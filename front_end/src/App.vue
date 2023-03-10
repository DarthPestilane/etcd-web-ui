<script setup>
import { SaveOutlined } from '@ant-design/icons-vue';
import { onMounted, reactive, ref } from 'vue';
import { httpClient } from './api/api.js';
import { message, notification } from 'ant-design-vue';
import AppEditor from './components/AppEditor.vue';

const loading = ref(false);

const endpoint = ref('127.0.0.1:2379');
const username = ref('');
const password = ref('');
const keyPrefix = ref('');
const etcdKeys = reactive([]);
const currEtcdKeyIdx = ref(-1);
const keyContent = ref('');
const keyTTL = ref(-1);

const langModes = reactive(['text', 'json', 'yaml']);
const usingLangMode = ref(langModes[0]);

const showCreateModal = ref(false);
const createForm = reactive({
  key: '',
  content: '',
});

onMounted(async () => {
  const resp = await httpClient.get('/options');
  console.log(resp.data);
  endpoint.value = resp.data.data.endpoint;
  username.value = resp.data.data.username;
  password.value = resp.data.data.password;
  keyPrefix.value = resp.data.data.keyPrefix;
});

const onRefresh = async () => {
  loading.value = true;

  try {
    const resp = await httpClient.post('/refresh', {
      endpoint: endpoint.value,
      username: username.value,
      password: password.value,
      keyPrefix: keyPrefix.value,
    });
    console.log(resp.data);
    etcdKeys.splice(0);
    resp.data.data.etcdKeys.forEach(v => {
      etcdKeys.push(v);
    });
    currEtcdKeyIdx.value = -1;
    keyContent.value = '';
  } catch (e) {
    handleApiErr(e);
  }

  loading.value = false;
};

const onGetEtcdKey = async (idx) => {
  loading.value = true;
  currEtcdKeyIdx.value = idx;
  try {
    const resp = await httpClient.post('/get', {
      endpoint: endpoint.value,
      username: username.value,
      password: password.value,
      etcdKey: etcdKeys[idx],
    });
    console.log(resp.data);
    keyContent.value = resp.data.data.content;
    keyTTL.value = resp.data.data.ttl;
  } catch (e) {
    handleApiErr(e);
  }

  loading.value = false;
};

const onKeyContentSave = async () => {
  loading.value = true;

  try {
    await httpClient.put('/put', {
      endpoint: endpoint.value,
      username: username.value,
      password: password.value,
      etcdKey: etcdKeys[currEtcdKeyIdx.value],
      content: keyContent.value,
    });
    await onGetEtcdKey(currEtcdKeyIdx.value);
  } catch (e) {
    handleApiErr(e);
  }

  loading.value = false;
};

const onKeyDelete = async (idx) => {
  loading.value = true;

  try {
    await httpClient.post('/delete', {
      endpoint: endpoint.value,
      username: username.value,
      password: password.value,
      etcdKey: etcdKeys[idx],
    });

    etcdKeys.splice(idx, 1);
    if (idx === currEtcdKeyIdx.value) {
      currEtcdKeyIdx.value = -1;
      keyContent.value = '';
      keyTTL.value = -1;
    }
  } catch (e) {
    handleApiErr(e);
  }

  loading.value = false;
};

const onCreateModalCancel = async () => {
  createForm.key = '';
  createForm.content = '';
};

const onCreateModalConfirm = async () => {
  if (!createForm.key) {
    message.error('key is required');
    return;
  }
  loading.value = true;

  try {
    await httpClient.put('/put', {
      endpoint: endpoint.value,
      username: username.value,
      password: password.value,
      etcdKey: createForm.key,
      content: createForm.content,
    });
  } catch (e) {
    handleApiErr(e);
  }

  showCreateModal.value = false;
  loading.value = false;

  await onRefresh();
};

const handleApiErr = (e) => {
  console.log(e);
  notification.error({
    message: e.message,
    description: `${e.response.data.err_msg} [${e.response.data.code}]`,
    duration: 0,
  });
};

</script>

<template>
  <a-layout>
    <a-spin :spinning="loading">
      <a-layout-content style="padding: 10px">
        <!--??????-->
        <h2>Connection</h2>
        <a-row>
          <a-form autocomplete="off" layout="inline">
            <a-form-item label="endpoint">
              <a-input v-model:value="endpoint" />
            </a-form-item>
            <a-form-item label="username">
              <a-input v-model:value="username" />
            </a-form-item>
            <a-form-item label="password">
              <a-input-password v-model:value="password" />
            </a-form-item>
            <a-form-item label="keyPrefix">
              <a-input v-model:value="keyPrefix" />
            </a-form-item>
            <a-form-item>
              <a-button type="primary" html-type="submit" @click="onRefresh">?????? keys</a-button>
            </a-form-item>
          </a-form>
        </a-row>

        <a-divider />

        <!--?????? modal ???-->
        <a-modal
          v-model:visible="showCreateModal"
          title="?????? etcd key value" width="1200px"
          ok-text="??????" cancel-text="??????"
          :afterClose="onCreateModalCancel"
          @ok="onCreateModalConfirm"
        >
          <a-form autocomplete="off" :labelCol="{span: 2}">
            <a-form-item label="key">
              <a-input v-model:value="createForm.key" />
            </a-form-item>
            <a-form-item label="value">
              <a-textarea v-model:value="createForm.content" :auto-size="{ minRows: 20 }" allowClear />
            </a-form-item>
          </a-form>
        </a-modal>

        <a-row>
          <!--?????????-->
          <a-col :span="10" class="etcdKeyList">
            <a-row style="border-bottom: 1px solid silver">
              <a-col>
                <h2>Etcd Keys</h2>
              </a-col>
              <a-col style="margin-left: 10px">
                <a-button type="primary" @click="() => showCreateModal = true">??????</a-button>
              </a-col>
            </a-row>
            <a-row
              v-for="(item, idx) in etcdKeys" :key="idx"
              class="etcdKeyItem"
              type="flex"
              align="middle"
            >
              <a-col :span="1">
                <span v-show="currEtcdKeyIdx===idx">*</span>
              </a-col>
              <a-col :span="21" @click="onGetEtcdKey(idx)">
                {{ item }}
              </a-col>
              <a-col :span="2">
                <a-popconfirm :title="`??????????????? ${etcdKeys[idx]} ?`" @confirm="onKeyDelete(idx)">
                  <a-button type="primary" danger size="small">??????</a-button>
                </a-popconfirm>
              </a-col>
            </a-row>
          </a-col>

          <!--??????-->
          <a-col :span="14">
            <a-row style="border-bottom: 1px solid silver">
              <h2>Content</h2>
            </a-row>
            <a-row>
              <a-col :span="24">
                <a-space align="center">
                  <a-select v-model:value="usingLangMode">
                    <a-select-option v-for="(item) in langModes" :key="item" :value="item">{{ item }}</a-select-option>
                  </a-select>
                  <a-popconfirm :title="`??????????????? ${etcdKeys[currEtcdKeyIdx]} ?`" @confirm="onKeyContentSave">
                    <a-button type="primary" v-show="currEtcdKeyIdx >= 0">
                      <template #icon>
                        <SaveOutlined />
                      </template>
                      ??????
                    </a-button>
                  </a-popconfirm>
                  <span>TTL: {{ keyTTL }}</span>
                </a-space>
              </a-col>
            </a-row>
            <a-row style="margin-top: 10px">
              <a-col :span="24">
                <AppEditor v-model="keyContent" :lang="usingLangMode" />
              </a-col>
            </a-row>
          </a-col>
        </a-row>
      </a-layout-content>
    </a-spin>
  </a-layout>
</template>

<style scoped lang="scss">
.etcdKeyList {
  border-right: 1px solid silver;

  .etcdKeyItem {
    cursor: pointer;
    margin: 2px 0;

    &:hover {
      background-color: silver;
    }
  }
}
</style>
