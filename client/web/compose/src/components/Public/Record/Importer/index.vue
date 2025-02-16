<template>
  <div class="d-flex">
    <b-button
      size="lg"
      variant="none"
      class="flex-fill"
      @click="showModal=true"
    >
      <font-awesome-icon :icon="['fas', 'file-import']" />
      {{ $t('general.label.import') }}
    </b-button>

    <b-modal
      :visible="showModal"
      size="lg"
      :title="$t('recordList.import.to', { modulename: module.name })"
      hide-footer
      ok-variant="primary"
      body-class="p-0"
      no-fade
      scrollable
      @hide="onModalHide"
    >
      <component
        :is="stepComponent"
        v-bind="$props"
        :session="session"
        @fileUploaded="onFileUploaded"
        @fieldsMatched="onFieldsMatched"
        @importSuccessful="onImportSuccessful"
        @importFailed="onImportFailed"
        @back="onBack"
        @reset="onReset"
        @close="onClose"
        v-on="$listeners"
      >
        <label
          v-if="progress.failed"
          slot="uploadLabel"
          class="text-danger"
        >

          {{ $t('recordList.import.failed', progress) }}
        </label>
      </component>
    </b-modal>
  </div>
</template>

<script>
import FileUpload from './FileUpload'
import FieldMatch from './FieldMatch'
import Progress from './Progress'
import ErrorReport from './ErrorReport'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'Importer',
  props: {
    namespace: {
      type: Object,
      required: true,
      default: () => ({}),
    },
    module: {
      type: Object,
      required: true,
      default: () => ({}),
    },
  },

  data () {
    return {
      step: 0,
      showModal: false,
      session: {},
      components: [FileUpload, FieldMatch, Progress, ErrorReport],
    }
  },

  computed: {
    stepComponent () {
      return this.components[this.step]
    },

    progress () {
      return this.session.progress || {}
    },
  },

  methods: {
    onModalHide () {
      this.step = 0
      this.session = {}
      this.showModal = false
    },

    onBack () {
      this.step = Math.max(0, this.step - 1)
    },

    onFileUploaded (e) {
      this.session = e
      this.step = 1
    },

    onFieldsMatched (e) {
      this.session.fields = e
      this.step = 2

      this.$ComposeAPI.recordImportRun(this.session)
    },

    onImportSuccessful () {
      this.$emit('importSuccessful')
    },

    onImportFailed (e) {
      this.session.progress = e
      this.step = 3
    },

    onReset () {
      this.step = 0
      this.$set(this, 'session', {})
      this.$emit('reset')
    },

    onClose () {
      this.showModal = false
    },
  },
}
</script>
