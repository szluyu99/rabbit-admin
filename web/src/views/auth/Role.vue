<script setup>
import { onMounted, ref } from 'vue'
import { UButton, UInput, UModal, UTable, UTag, UTree } from 'unocss-ui'

import Date from '@/components/Date.vue'

import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import api from '@/api'

const permissionOptions = ref([])

const {
  list, keyword, loading, modalVisible, form, page, limit, total,
  handleQuery, handleSearch,
  handleShowAdd, handleAdd,
  handleShowEdit, handleEdit,
  handleDelete,
} = useTable({
  initForm: {
    name: '',
    label: '',
    permission_ids: [],
  },
  queryFn: params => request.post('/api/role', params),
  addFn: item => request.put('/api/role', item),
  deleteFn: id => request.delete(`/api/role/${id}`),
  editFn: item => request.patch(`/api/role/${item.id}`, item),
  afterQuery: () => {
    list.value = list.value.map(item => ({
      ...item,
      permission_ids: item.permissions.map(permission => permission.id),
    }))
  },
})

async function handleSaveOrUpdate() {
  try {
    form.value.id ? await handleEdit(form.value) : await handleAdd()
    handleQuery()
  }
  catch (err) {
    console.error(err)
  }
}

async function handleInitDefaultRoles() {
  try {
    await api.initDefaultRoles()
    handleQuery()
  }
  catch (err) {
    console.error(err)
  }
}

onMounted(async () => {
  handleQuery()
  permissionOptions.value = await api.getPermissionOptions()
})
</script>

<template>
  <PageWrap title="Role Page">
    <template #description>
      This page can only be managed (added, edited, deleted) by <span class="font-bold"> super user </span>.
    </template>
    <div class="flex items-center justify-between">
      <div class="w-36 sm:w-52">
        <UInput
          v-model="keyword"
          right-icon="i-mdi:magnify"
          @click-right="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
      <div class="flex gap-2">
        <UButton type="error" @click="handleInitDefaultRoles">
          Init
        </UButton>
        <UButton type="primary" @click="handleShowAdd">
          Add Role
        </UButton>
      </div>
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" header-color>
        <template #headers>
          <th scope="col" class="table-th">
            Name
          </th>
          <th scope="col" class="table-th">
            Label
          </th>
          <th scope="col" class="table-th">
            CreatedAt
          </th>
          <th scope="col" class="table-th" />
        </template>
        <template #rows="{ row }">
          <td class="table-td">
            {{ row.name }}
          </td>
          <td class="table-td">
            <UTag type="success">
              {{ row.label }}
            </UTag>
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="row.createdAt" />
          </td>
          <td class="table-td">
            <div class="flex items-center gap-4 text-xl">
              <button
                class="i-mdi:delete text-red-400 hover:text-red-500"
                @click="handleDelete(row.id)"
              />
              <button
                class="i-mdi:book-edit text-green-400 hover:text-green-500"
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
      {{ form.id ? 'Edit Role' : 'Add Role' }}
    </template>
    <div class="space-y-3">
      <UInput v-model="form.name" label="Name" />
      <UInput v-model="form.label" label="Label" />
      <div class="space-y-2">
        <span> Role Permissions: </span>
        <UTree
          v-model="form.permission_ids"
          :options="permissionOptions"
          selectable
          cascade
        />
      </div>
    </div>
    <template #footer>
      <UButton v-if="form.id" type="warning" @click="handleSaveOrUpdate">
        Edit
      </UButton>
      <UButton v-else type="primary" @click="handleSaveOrUpdate">
        Add
      </UButton>
      <UButton @click="modalVisible = false">
        Cancel
      </UButton>
    </template>
  </UModal>
</template>
