<template>
  <b-input-group
    :size="size"
    class="input-group-border border-light rounded"
    style="min-width: 200px;"
  >
    <b-input
      ref="searchInput"
      data-test-id="input-search"
      :type="inputType"
      name="search"
      :value="value"
      :debounce="debounce"
      :disabled="disabled"
      :placeholder="placeholder"
      :autocomplete="autocomplete"
      class="h-100 pr-0 border-0 rounded-0 text-truncate bg-white"
      @input="onInput"
      @update="search"
      @keyup.enter="submitQuery"
    />
    <b-input-group-append
      v-if="showSubmittable"
      :class="{ 'search-icon-border border-light': showSubmittableAndClearable }"
      class="bg-white m-0"
    >
      <b-button
        :variant="isSubmittable ? 'outline-light' : 'link'"
        :disabled="disabled"
        :class="{
          'cursor-default': !isSubmittable
        }"
        class="d-inline-flex align-items-center border-0 rounded-0 py-0"
        @[isSubmittable]="submitQuery"
      >
        <font-awesome-icon
          :icon="['fas', 'search']"
          class="align-middle text-primary"
        />
      </b-button>
    </b-input-group-append>
  </b-input-group>
</template>

<script>
export default {
  name: 'CInputSearch',

  props: {
    value: {
      type: String,
      default: '',
    },

    placeholder: {
      type: String,
      default: '',
    },

    size: {
      type: String,
    },

    disabled: {
      type: Boolean,
    },

    clearable: {
      type: Boolean,
      default: true,
    },

    submittable: {
      type: Boolean,
      default: false,
    },

    autocomplete: {
      type: String,
      default: 'on',
    },

    debounce: {
      type: Number,
      default: 0,
    },
  },

  data () {
    return {
      localValue: this.value,
    }
  },

  computed: {
    inputType () {
      return this.clearable ? 'search' : 'text'
    },

    showSubmittable () {
      return !this.localValue || this.showSubmittableAndClearable
    },

    isSubmittable () {
      return this.submittable && !this.disabled ? 'click' : null
    },

    showSubmittableAndClearable () {
      return this.clearable && this.submittable
    },
  },

  methods: {
    onInput (value) {
      this.localValue = value
    },

    search (e) {
      if (!this.submittable) {
        this.$emit('input', e)
      }
    },

    submitQuery () {
      if (this.submittable) {
        this.$emit('search', this.$refs.searchInput.localValue)
      }
    },
  },
}
</script>

<style lang="scss" scoped>
input:focus::placeholder {
  color: transparent;
}

input[type="search"]::-webkit-search-cancel-button {
  -webkit-appearance: none;
  height: 1em;
  width: 1em;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg viewPort='0 0 12 12' version='1.1' xmlns='http://www.w3.org/2000/svg'%3e%3cline x1='1' y1='11' x2='11' y2='1' stroke='black' stroke-width='2'/%3e%3cline x1='1' y1='1' x2='11' y2='11' stroke='black' stroke-width='2'/%3e%3c/svg%3e");
  cursor: pointer;
  margin-right: 13px;
  margin-top: 2px;
}

.input-group-border {
  border-width: 1px;
  border-style: solid;
}

.search-icon-border {
  border-width: 0px 0px 0px 2px;
  border-style: solid;
}

.cursor-default {
  cursor: default !important;
}
</style>
