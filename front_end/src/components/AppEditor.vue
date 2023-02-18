<script setup>
import { VAceEditor } from 'vue3-ace-editor';
import ace from 'ace-builds';
import 'ace-builds/src-noconflict/theme-solarized_light';
import 'ace-builds/src-noconflict/mode-text';
import 'ace-builds/src-noconflict/mode-yaml';
import 'ace-builds/src-noconflict/mode-json';
import { computed } from 'vue';

ace.config.set('workerPath', '/ace');

const emit = defineEmits(['update:modelValue']);
const props = defineProps({
  modelValue: {
    type: String,
    required: true,
  },
  lang: {
    type: String,
    default: 'text',
  },
  readonly: {
    type: Boolean,
    default: false,
  },
});

const value = computed({
  get() {
    return props.modelValue;
  },
  set(value) {
    emit('update:modelValue', value);
  },
});

const options = {
  fontSize: 14,
  useWorker: true,
  fixedWidthGutter: true,
  showLineNumbers: true,
  displayIndentGuides: true,
};
</script>

<template>
  <v-ace-editor
    class="app-editor"
    v-model:value="value"
    :min-lines="10"
    :max-lines="80"
    theme="solarized_light"
    :lang="props.lang"
    :readonly="props.readonly"
    :options="options"
  />
</template>

<style scoped lang="scss">
.app-editor {
  line-height: 1.2;
  border: 1px solid silver;
  //box-shadow: 0 0 10px 2px silver;
}
</style>
