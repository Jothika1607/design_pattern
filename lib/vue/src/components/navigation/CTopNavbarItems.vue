<template>
  <div class="nav-sidebar">
    <b-button
      v-for="({page = {}, params = {}, children = []}) of items"
      :key="pageIndex(page)"
      variant="link"
      class="w-100 text-dark text-decoration-none p-0 nav-item d-flex"
      active-class="nav-active"
      exact-active-class="nav-active"
      :title="page.title"
      :to="{ name: page.name || defaultRouteName, params }"
    >
      <span
        class="d-inline-block text-nowrap text-truncate"
        @click="closeSidebar()"
      >
        <template
          v-if="page.icon"
        >
        <!-- <font-awesome-icon :icon="['fas', 'users']" /> -->
          <font-awesome-icon
            v-if="Array.isArray(page.icon)"
            class="icon"
            :icon="page.icon"
          />
          <template v-else>
            <img
              :src="page.icon"
              class="mr-1"
              style="height: 40px; width: 40px;"
            />
          </template>
        </template>
        <label
          class="title mb-0 pointer"
        >
          {{ page.title }}
        </label>
      </span>
 
      <template
        v-if="children.length"
      >
      <b-button
      variant="None"
      size="sm"
      class="text-primary p-0 border-0 float-right mr-1 caret-button"
      @click.self.stop.prevent="toggle(page)"
    >
      <font-awesome-icon
        v-if="!collapses[pageIndex(page)]"
        class="pointer-none"
        :icon="['fas', 'caret-down']"
      />
      <font-awesome-icon
        v-else
        class="pointer-none"
        :icon="['fas', 'caret-up']"
      />
    </b-button>

    <b-collapse
      class="position-absolute sub-nav-item"
      :visible="collapses[pageIndex(page)]"
      @click.stop.prevent
      ref="dropdown"
    >
          <c-top-navbar-items
            :items="children"
            :start-expanded="startExpanded"
            :default-route-name="defaultRouteName"
            v-on="$listeners"
          />
        </b-collapse>
      </template>
    </b-button>
  </div>
</template>
 
<script>
export default {
  name: 'CTopNavbarItems',
 
  props: {
    /*
    * {
        page: { name, title }
        params: {...}
      }
    */
    items: {
      type: Array,
      required: true,
      default: () => [],
    },
    defaultRouteName: {
      type: String,
      required: true,
    },
    startExpanded: {
      type: Boolean,
      required: false,
    },
  },
 
  data () {
    return {
      collapses: {},
    }
  },
 
  watch: {
    items: {
      immediate: true,
      handler (items = []) {
        items.forEach(({ page, params, children }) => {
          const px = this.pageIndex(page)
          // Apply startExpanded only if page isn't currently expanded
          this.$set(this.collapses, px, this.startExpanded || page.expanded || this.showChildren({ params, children }))
        })
      },
    },
  },
 
  methods: {
    closeSidebar () {
      if (window.innerWidth < 576) {
        this.$root.$emit('close-sidebar')
      }
    },
 
    toggle(p) {
      const px = this.pageIndex(p);
      this.$set(this.collapses, px, !this.collapses[px]);
    },
    pageIndex(p) {
      return p.pageID || p.name || p.title;
    },
    closeDropdownOnOutsideClick(event) {
      // Check if the clicked element is outside the dropdown
      if (
        this.$refs.dropdown &&
        !this.$refs.dropdown.$el.contains(event.target)
      ) {
        // Close the dropdown
        this.collapses.forEach((_, index) => {
          this.collapses[index] = false;
        });
      }
    },
 
    // Recursively check for child pages that are open, so that parents can open as well
    showChildren ({ params = {}, children = [] }) {
      const partialParamsMatch = Object.entries(params).some(([key, value]) => {
        return this.$route.params[key] === value
      })
 
      if (partialParamsMatch) {
        return partialParamsMatch
      }
 
      return children.map(c => this.showChildren(c)).some(isOpen => isOpen)
    },
  },

  mounted() {
    // Add a global click event listener
    document.addEventListener("click", this.closeDropdownOnOutsideClick);
  },
  beforeDestroy() {
    // Remove the global click event listener to prevent memory leaks
    document.removeEventListener("click", this.closeDropdownOnOutsideClick);
  },
}
</script>
 
<style scoped lang="scss">
// This has to be there, so chevrons are clickable inside the button
.pointer-none {
  pointer-events: none;
}
.caret-button{
  width:40px;
  height: 40px;
  margin-left: 8px;
}
 
.svg-inline--fa {
  width: 30px;
}
 
.sub-nav-item{
  position: absolute;
  transform: translateY(81px);
  z-index: 2;
  >div{
    background: white;
    border: 1px solid white;
    flex-direction: column;
  }
  
  .nav-sidebar {
    width: 100% !important;
    padding-left: 0px !important;
    padding-right: 0px !important;
    transform: translateX(-21px);
  }
 
  .nav-active{
    border-bottom-color: white !important;
  }
}
 
.nav-item > span {
  .title {
    color: var(--tertiary);
  }
}
 
.nav-item:hover > span {
  .title {
    color: var(--primary);
    transition: color 0.25s;
  }
}
 
.nav-active > span > {
  .icon {
    color: var(--primary)
  }
 
  .title {
    font-family: 'Roboto-Regular';
    color: var(--primary);
    transition: color 0.5s
  }
}
 
.nav-item {
  text-align: left;
}
 
[dir="rtl"] {
  .nav-item {
    text-align: right;
  }
}
</style>