<template>
  <div>
    <div class="uk-flex uk-flex-column" id="accounts-app">
      <template v-if="isInitialized">
        <div class="oc-app-bar">
          <accounts-batch-actions
            v-if="isAnyAccountSelected"
            :number-of-selected-accounts="numberOfSelectedAccounts"
            :selected-accounts="selectedAccounts"
          />
          <accounts-create v-else />
        </div>
        <oc-grid class="uk-height-1-1 uk-flex-1 uk-overflow-auto">
          <div class="uk-width-expand">
            <accounts-list :accounts="accounts" />
          </div>
        </oc-grid>
      </template>
      <oc-loader v-else />
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions, mapState } from 'vuex'
import AccountsList from './accounts/AccountsList.vue'
import AccountsCreate from './accounts/AccountsCreate.vue'
import AccountsBatchActions from './accounts/AccountsBatchActions.vue'

export default {
  name: 'App',
  components: { AccountsBatchActions, AccountsList, AccountsCreate },
  computed: {
    ...mapGetters('Accounts', ['isInitialized', 'getAccountsSorted', 'isAnyAccountSelected']),
    ...mapState('Accounts', ['selectedAccounts']),

    accounts () {
      return this.getAccountsSorted
    },
    numberOfSelectedAccounts () {
      return this.selectedAccounts.length
    }
  },
  methods: {
    ...mapActions('Accounts', ['initialize'])
  },
  created () {
    this.initialize()
  }
}
</script>

<style>
/* TODO: After https://github.com/owncloud/owncloud-design-system/pull/418 gets merged
there won't be an extra span and this won't be needed anymore */
.accounts-selection-actions-btn > span {
  display: flex;
  align-items: center;
}

/* TODO: Adjust in ODS */
.oc-dropdown-menu {
  width: 150px;
}
</style>
