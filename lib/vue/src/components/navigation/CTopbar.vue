<template>
  <div class="header-navigation d-flex flex-wrap align-items-center pr-3">
    <div
      class="spacer"
      :class="{
        'expanded': sidebarPinned,
      }"
    />
    <!-- <h2 class="title py-1 mb-0">
      <slot name="title" />
    </h2> -->
    <!-- <div>
      <img src="" alt="DataSirpi"/>
    </div> -->
    <div class="icon-logo">
      <img
      width="78px"
      height="32px"
      src="
      http://52.66.61.197:18080/assets/logo.png"/>
    </div>

    <div class="search-container">
      <input
              v-model="searchTerm"
              @input="handleInput"
              type="text"
              placeholder="Search case here"
              class="search-bar"
            />
      <span class="search-icon">
        <font-awesome-icon :icon="['fas', 'search']" />
      </span>
      </div>   
 
    <div class="tools-wrapper py-1 ml-auto mr-4">
      <slot name="tools" />
    </div>
 
    <div class="d-flex align-items-center ml-auto icon-space">
        <div class="d-flex icon-colour align-items-center">
      <b-button
        v-if="!hideAppSelector && !settings.hideAppSelector"
        data-test-id="app-selector"
        variant="outline-light"
        title="Apps"
        size="lg"
        :href="appSelectorURL"
        class="d-flex align-items-center justify-content-center border-0 nav-icon rounded-circle text-sm-nowrap icon-position"
      >
        <font-awesome-icon
          class="m-0 h5 icon-grip"
          :icon="['fas', 'grip-horizontal']"
        /> 
      

      </b-button>
          <b-dropdown
             text="Case Management"
             variant="none"
             class="m-md-2 dropdown-style"
           >
               <b-dropdown-item>Case Management</b-dropdown-item>
               <b-dropdown-item href="http://52.66.61.197:18080/admin/system/application/list">Admin View</b-dropdown-item>
           </b-dropdown>
         </div>

         <b-button
         v-if="!hideAppSelector && !settings.hideAppSelector"
         data-test-id="app-selector"
         variant="none"
         size="lg"
         class="d-flex align-items-center justify-content-center text-dark border-0 nav-icon rounded-circle text-sm-nowrap"
       >
         <font-awesome-icon
           class="m-0 h5 icon-bell" 
           :icon="['far', 'bell']"
         />
       </b-button>
 
      <b-dropdown
        v-if="!settings.hideHelp"
        data-test-id="dropdown-helper"
        size="lg"
        variant="outline-light"
        toggle-class="text-decoration-none text-dark rounded-circle border-0 w-100"
        menu-class="topbar-dropdown-menu border-0 shadow-sm text-dark font-weight-bold mt-2"
        right
        no-caret
        class="nav-icon mx-1 text-sm-nowrap"
      >
        <template #button-content>
          <div
            class="d-flex align-items-center justify-content-center"
          >
            <font-awesome-icon
              class="m-0 h5 icon-question"
              :icon="['fas', 'question-circle']"
            />
            <span class="sr-only">
              {{ labels.helpForum }}
            </span>
          </div>
        </template>
        <div>
          <slot name="help-dropdown" />
        </div>
        <b-dropdown-item
          v-for="(helpLink, index) in helpLinks"
          :key="index"
          :href="helpLink.url | checkValidURL"
          :target="helpLink.newTab ? '_blank' : ''"
        >
          {{ helpLink.handle }}
        </b-dropdown-item>
        <b-dropdown-item
          v-if="!settings.hideForumLink"
          data-test-id="dropdown-helper-forum"
          href="https://forum.cortezaproject.org/"
          target="_blank"
        >
          {{ labels.helpForum }}
        </b-dropdown-item>
        <b-dropdown-item
          v-if="!settings.hideDocumentationLink"
          data-test-id="dropdown-helper-docs"
          :href="documentationURL"
          target="_blank"
        >
          {{ labels.helpDocumentation }}
        </b-dropdown-item>
        <b-dropdown-item
          v-if="!settings.hideFeedbackLink"
          data-test-id="dropdown-helper-feedback"
          href="mailto:info@cortezaproject.org"
          target="_blank"
        >
          {{ labels.helpFeedback }}
        </b-dropdown-item>
        <b-dropdown-divider
          v-if="!onlyVersion"
        />
        <b-dropdown-item
          disabled
        >
          {{ labels.helpVersion }}
          <br>
          {{ frontendVersion }}
        </b-dropdown-item>
      </b-dropdown>
 
      <b-dropdown
        v-if="!settings.hideProfile"
        data-test-id="dropdown-profile"
        data-v-onboarding="profile"
        :variant="avatarExists ? 'link' : 'outline-light'"
        :toggle-class="`nav-icon text-decoration-none rounded-circle border d-flex align-items-center carat-test  p-0  ${avatarExists ? 'p-0' : ''}`"
        right
        menu-class="topbar-dropdown-menu border-0 shadow-sm font-weight-bold mt-2 "
        class="nav-user-icon profile"
      >
        <template #button-content>
          <div
            v-if="avatarExists"
            class="avatar d-flex"
            :style="{
              'background-image': avatarExists  ? `url(${profileAvatarUrl})` : 'none',
            }"
          />
 
          <div
            v-else
            class="d-flex align-items-center justify-content-center"
          >
            <font-awesome-icon
              class="m-0 h5 icon-profile"
              :icon="['far', 'user']"
            />
            <span class="sr-only">
              {{ labels.helpForum }}
            </span>
          </div>
        </template>
        <b-dropdown-text
          data-test-id="dropdown-item-username"
          class="text-muted mb-2"
        >
          {{ labels.userSettingsLoggedInAs }}
        </b-dropdown-text>
        <div>
          <slot name="avatar-dropdown" />
        </div>
        <b-dropdown-item
          v-for="(profileLink, index) in profileLinks"
          :key="index"
          :href="profileLink.url | checkValidURL"
          :target="profileLink.newTab ? '_blank' : ''"
        >
          {{ profileLink.handle }}
        </b-dropdown-item>
        <b-dropdown-item
          v-if="!settings.hideProfileLink"
          data-test-id="dropdown-profile-user"
          :href="userProfileURL"
          target="_blank"
        >
          {{ labels.userSettingsProfile }}
        </b-dropdown-item>
        <b-dropdown-item
          v-if="!settings.hideChangePasswordLink"
          data-test-id="dropdown-profile-change-password"
          :href="changePasswordURL"
          target="_blank"
        >
          {{ labels.userSettingsChangePassword }}
        </b-dropdown-item>
        <b-dropdown-divider />
        <b-dropdown-item
          data-test-id="dropdown-profile-logout"
          href=""
          @click="$auth.logout()"
          class="mt-2"
        >
          {{ labels.userSettingsLogout }}
        </b-dropdown-item>
      </b-dropdown>
    </div>
    <div class="w-100">
     <slot name="top-nav-bar"/>
    </div>
  </div>
</template>
 
<script>
export default {
  props: {
    sidebarPinned: {
      type: Boolean,
      required: true,
      default: false,
    },
 
    hideAppSelector: {
      type: Boolean,
      required: false,
      default: false,
    },
 
    appSelectorURL: {
      type: String,
      default: '../'
    },
 
    settings: {
      type: Object,
      required: true,
    },
 
    labels: {
      type: Object,
      required: true,
    },
  },
 
  computed: {
    userProfileURL () {
      return this.$auth.cortezaAuthURL
    },
 
    changePasswordURL () {
      return `${this.$auth.cortezaAuthURL}/change-password`
    },
 
    documentationURL () {
      const [year, month] = VERSION.split('.')
      return `https://docs.cortezaproject.org/corteza-docs/${year}.${month}/index.html`
    },
 
    helpLinks () {
      const { helpLinks = [] } = this.settings || {}
      return (helpLinks || []).filter(({ handle, url }) => handle && url)
    },
 
    profileLinks () {
      const { profileLinks = [] } = this.settings || {}
      return (profileLinks || []).filter(({ handle, url }) => handle && url)
    },
 
    onlyVersion () {
      const {
        hideForumLink,
        hideDocumentationLink,
        hideFeedbackLink,
      } = this.settings || {}
 
      return !this.helpLinks.length && hideForumLink && hideDocumentationLink && hideFeedbackLink
    },
 
    frontendVersion () {
      /* eslint-disable no-undef */
      return VERSION
    },
 
    profileAvatarUrl () {
      return `${this.$SystemAPI.baseURL}/attachment/avatar/${this.$auth.user.meta.avatarID}/original/profile-photo-avatar`
    },
 
    avatarExists () {
      return this.$auth.user.meta.avatarID !== "0" && this.$auth.user.meta.avatarID
    },
  }
}
</script>
 
<style lang="scss" scoped>
$header-height: 96px;
$nav-width: 1440px;
$nav-icon-size: 40px;
$nav-user-icon-size: 40px;
 
.icon-logo {
  width: 126px;
  height: 96px;
  display: flex;
  align-items: center;
  justify-content: center;
}
 
.nav-icon {
  width: $nav-icon-size;
  height: $nav-icon-size;
}
 
.nav-user-icon {
  width: 40px;
  height: 40px;
}
 
.header-navigation {
  width: 100vw;
  min-height: $header-height;
  background-color: var(--body-bg);
}
 
.avatar {
  border-radius: 50%;
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center;
  width: 40px;
  height: 40px;
}
 
.spacer {
  min-width: 0px;
  -webkit-transition: min-width 0.15s ease-in-out;
  -moz-transition: min-width 0.15s ease-in-out;
  -o-transition: min-width 0.15s ease-in-out;
  transition: min-width 0.15s ease-in-out;
 
  &.expanded {
    min-width: calc(#{$nav-width} - 42px);
    -webkit-transition: min-width 0.2s ease-in-out;
    -moz-transition: min-width 0.2s ease-in-out;
    -o-transition: min-width 0.2s ease-in-out;
    transition: min-width 0.2s ease-in-out;
  }
}
 
.title {
  display: flex;
  align-items: center;
  min-height: $header-height;
  padding-left: calc(3.5rem + 6px);
 
  .vue-portal-target {
    display: -webkit-box; /* For Safari and old versions of Chrome */
    display: -ms-flexbox; /* For old versions of IE */
    -webkit-box-orient: vertical; /* For Safari and old versions of Chrome */
    -webkit-line-clamp: 3; /* Maximum number of lines to display */
    overflow: hidden;
    text-overflow: ellipsis;
  }
}
 
.tools-wrapper {
  flex-grow: 1;
 
  .vue-portal-target {
    display: flex;
    justify-content: end;
    align-items: center;
    flex-wrap: wrap;
  }
}

.icon-space{
  gap: 56px;
  width: 626px;
  height: 40px;
}
.icon-colour :hover{
  color: #3D6ECF;
}

.search-container {
  position: relative;
  display: inline-block;
  margin-left: 40px;
  margin-top: 24px;
  margin-bottom: 24px;
  border: none;
}
 
.search-icon {
  position: absolute;
  top: 50%;
  right: 20px;
  width: 32px;
  height: 32px;
  transform: translateY(-50%);
  color:#373737;
  pointer-events: none;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-bar{
    width: 497px;
    height: 48px;
    border: 1px solid #878787;
    border-radius: 8px;
    padding-left: 20px;
    color: #878787 !important;
}

span.search-icon svg.fa-search{
  width:23px;
  height:23px;
}

.search-bar::placeholder {
width: 145px;
height:22px; /* Set the width of the placeholder text */
line-height: 22px; /* Add padding to the left of the placeholder text */
}

.icon-grip{
  width: 28px !important;
  height: 28px !important;
  color: #525252;
}

.icon-bell{
  width: 28px !important;
  height: 32px !important;
  color: #525252;
}

.icon-question{
  width: 32px !important;
  height: 32px !important;
  color: #525252;
}
.icon-profile{
  width: 40px !important;
  height: 32px !important;
  color: #525252;
}

.profile{
  width: 72px;
  height: 40px;

}

</style>
 
<style lang="scss">
.topbar-dropdown-menu {
  z-index: 1100;
}
</style>
