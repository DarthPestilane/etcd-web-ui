<script setup>
import { onMounted, watch } from 'vue';
import ace from 'ace-builds';
import 'ace-builds/src-noconflict/theme-solarized_light';
import 'ace-builds/src-noconflict/mode-yaml';
import 'ace-builds/src-noconflict/mode-json';
import 'ace-builds/src-noconflict/mode-text';

/** @type Ace.Editor */
let editor = null;

const props = defineProps(['modelValue', 'langMode']);
const emit = defineEmits(['update:modelValue']);

ace.config.set('workerPath', '/ace');

onMounted(async () => {
  editor = ace.edit('ace-editor', {
    maxLines: 80,
    minLines: 40,
    fontSize: 14,
    showLineNumbers: true,
    useWorker: true,
    displayIndentGuides: true,
    theme: 'ace/theme/solarized_light',
    mode: `ace/mode/${props.langMode}`,
  });

  editor.on('change', () => {
    emit('update:modelValue', editor.getValue());
  });
});

watch(
  () => props.modelValue,
  (newVal) => {
    const position = editor.getCursorPosition();
    editor.getSession().setValue(newVal);
    editor.clearSelection();
    editor.moveCursorToPosition(position);
  },
);

watch(
  () => props.langMode,
  (newVal) => {
    editor.getSession().setMode(`ace/mode/${newVal}`);
  },
);
</script>

<template>
  <div class="ace-editor" id="ace-editor"></div>
</template>

<style scoped lang="scss">
</style>
