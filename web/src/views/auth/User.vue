<script setup>
import { onMounted, ref } from 'vue'
import { UButton, UCheckbox, UInput, UModal, UTable } from 'unocss-ui'

import Date from '@/components/Date.vue'
import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'

const {
  list, keyword, loading, page, limit, total, form, modalVisible,
  handleQuery, handleSearch, handleShowEdit,
} = useTable({
  queryFn: params => request.post('/api/user', params),
  addFn: item => request.put('/api/user', item),
  deleteFn: id => request.delete(`/api/user/${id}`),
  editFn: item => request.patch(`/api/user/${item.id}`, item),
  batchFn: ids => request.delete('/api/user', ids),
  initForm: {
    role_ids: [],
  },
})

async function handleQueryList() {
  await handleQuery()
  list.value = list.value.map(item => ({
    ...item,
    role_ids: item.roles.map(role => role.id),
  }))
}

function handleUpdateUserRole() {
  request.patch('/api/user/role', form.value).then(() => {
    modalVisible.value = false
    handleQueryList()
  })
}

const roleOptions = ref([])

onMounted(() => {
  request.post('/api/role', { page: 1, limit: 150 }).then((res) => {
    roleOptions.value = (res.items || []).map(item => ({
      label: item.label,
      value: item.id,
    }))
  })

  handleQueryList()
})
</script>

<template>
  <PageWrap title="User Page">
    <template #description>
      You can manage users here.
    </template>
    <div class="flex items-center justify-between">
      <div class="w-52">
        <UInput
          v-model="keyword"
          right-icon="i-mdi:magnify"
          @click-right="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" header-color>
        <template #headers>
          <th scope="col" class="table-th">
            Email
          </th>
          <th scope="col" class="table-th">
            Enabled
          </th>
          <th scope="col" class="table-th">
            LastLogin
          </th>
          <th scope="col" class="table-th">
            LastLoginIP
          </th>
          <th scope="col" class="table-th" />
        </template>
        <template #rows="{ row }">
          <td class="table-td">
            {{ row.email }}
          </td>
          <td class="table-td flex items-center">
            <div v-if="row.enabled" class="i-heroicons:check-circle-solid inline-block h-6 w-6 text-green-500" />
            <div v-else class="i-heroicons:x-circle-solid inline-block h-6 w-6 text-red-500" />
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="row.lastLogin" />
          </td>
          <td class="table-td whitespace-nowrap">
            {{ row.lastLoginIP }}
          </td>
          <td class="table-td">
            <div class="flex items-center gap-2 text-xl">
              <span
                class="i-mdi:human-edit cursor-pointer text-green-400 hover:text-green-500"
                @click="handleShowEdit(row)"
              />
            </div>
          </td>
        </template>
      </UTable>
    </div>
    <Pagination
      v-model:page="page"
      v-model:limit="limit"
      :total="total"
      @query="handleQuery"
    />
  </PageWrap>

  <UModal v-model="modalVisible">
    <template #header>
      Edit User
    </template>
    <div class="space-y-3">
      <p> User Role </p>
      <div class="flex gap-4">
        <UCheckbox
          v-for="item of roleOptions" :key="item.value"
          v-model="form.role_ids"
          :value="item.value"
        >
          {{ item.label }}
        </UCheckbox>
      </div>
    </div>
    <template #footer>
      <UButton v-if="form.id" type="warning" @click="handleUpdateUserRole">
        EDIT
      </UButton>
      <UButton @click="modalVisible = false">
        CANCEL
      </UButton>
    </template>
  </UModal>
</template>
