<template>
  <div class="h-100">
    <b-card
      no-body
      class="d-flex flex-column h-100 shadow overflow-hidden position-static"
      :class="[blockClass, cardClass]"
    >
      <b-card-header
        v-if="showHeader"
        class="border-0 text-nowrap"
        header-bg-variant="white"
        :header-text-variant="block.style.variants.headerText"
      >
        <div v-if="!headerSet">
          <div class="d-flex justify-content-between">
            <h5
              v-if="blockTitle"
              class="text-truncate mb-0 headtitle"
            >
              {{ blockTitle }}

              <slot name="title-badge" />
            </h5>
            <div v-if="toolbarSet">
              <slot name="toolbar" />
            </div>
          </div>
          <div class="d-flex justify-content-between">
            <b-button-group
              v-if="showOptions"
              class="shortcut-icon"
            >
              <b-button
                v-if="block.options.showRefresh"
                v-b-tooltip.hover="{
                  title: $t('general.label.refresh'),
                  container: '#body',
                }"
                variant="none"
                class="d-flex align-items-center shortcuticon-color d-print-none border-0"
                @click="$emit('refreshBlock')"
              >
                <font-awesome-icon
                  :icon="['fa', 'sync']"
                  class="shortcuticon-medium"
                />
              </b-button>
              <b-button
                v-if="block.options.magnifyOption || isBlockMagnified"
                v-b-tooltip.hover="{
                  title: isBlockMagnified ? '' : $t('general.label.magnify'),
                  container: '#body',
                }"
                variant="none"
                class="d-flex align-items-center shortcuticon-color d-print-none border-0"
                @click="
                  $root.$emit(
                    'magnify-page-block',
                    isBlockMagnified ? undefined : magnifyParams
                  )
                "
              >
                <font-awesome-icon
                  :icon="['fas', isBlockMagnified ? 'times' : 'search-plus']"
                  class="shortcuticon-medium"
                />
              </b-button>
            </b-button-group>
            <div>
              <slot name="searchcase" />
            </div>
          </div>
        </div>
        <b-card-text
          v-if="blockDescription"
          class="text-dark text-wrap mt-1"
        >
          {{ blockDescription }}
        </b-card-text>
        <slot
          v-else
          name="header"
        />
      </b-card-header>
      <b-card-body
        class="p-0 flex-fill"
        :class="{ 'overflow-auto': scrollableBody }"
      >
        <slot name="default" />
      </b-card-body>
      <b-card-footer
        v-if="footerSet"
        class="p-0 bg-white border-top"
      >
        <slot name="footer" />
      </b-card-footer>
    </b-card>
  </div>
</template>
<script>
import base from './base.vue'
export default {
  name: 'CardWrap',
  extends: base,
}
</script>
