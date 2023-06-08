<script setup>
import { onMounted, ref } from 'vue'
import { UButton, UDivider, UInput, UModal, URadio, USwitch, UTable, UTag } from 'unocss-ui'

import Date from '@/components/Date.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import api from '@/api'

const initForm = ref({
  name: '',
  uri: '',
  method: 'GET',
  parentId: 0,
})

const {
  list, keyword, loading, modalVisible, form,
  handleQuery, handleSearch,
  handleShowAdd, handleAdd,
  handleShowEdit, handleEdit,
  handleDelete, handleBatchDelete,
} = useTable({
  initForm: initForm.value,
  queryFn: params => request.post('/api/permission', params),
  addFn: item => request.put('/api/permission', item),
  deleteFn: id => request.delete(`/api/permission/${id}`),
  editFn: item => request.patch(`/api/permission/${item.id}`, item),
})

const {
  modalVisible: moduleModalVisible,
  form: moduleForm,
  handleShowAdd: handleShowAddModule,
  handleAdd: handleAddModule,
  handleShowEdit: handleShowEditModule,
  handleEdit: handleEditModule,
} = useTable({
  initForm: {
    objectName: '',
    name: '',
    parentId: 0,
  },
  queryFn: handleQuery,
  addFn: item => request.put('/api/permission', item),
  editFn: item => request.patch(`/api/permission/${item.id}`, item),
  deleteFn: id => request.delete(`/api/permission/${id}`),
})

const bulkActions = [
  {
    name: 'Delete All',
    clicked: keys => handleBatchDelete(keys),
  },
]

async function handleInitDefaultPermissions() {
  await api.initDefaultPermission()
  handleQuery()
}

async function handleCreateWebobjectPermissions() {
  await api.createWebobjectPermissions(moduleForm.value.objectName)
  moduleModalVisible.value = false
  handleQuery()
}

async function handleToggleAnonymous(row, val) {
  try {
    row.toggleLoading = true
    await request.patch(`/api/permission/${row.id}`, { anonymous: val })
    setTimeout(() => {
      row.toggleLoading = false
      row.anonymous = val
    }, 300)
  }
  catch (err) {
    row.toggleLoading = false
  }
}

onMounted(() => {
  handleQuery()
})

const methods = ['GET', 'POST', 'PUT', 'DELETE', 'PATCH']
const methodTypes = {
  GET: 'info',
  POST: 'success',
  PUT: 'warning',
  DELETE: 'error',
  PATCH: 'primary',
}
</script>

<template>
  <PageWrap title="Permission Page">
    <template #description>
      This page can only be managed (added, edited, deleted) by <span class="font-bold"> super user </span>.
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

      <div class="flex gap-2">
        <UButton type="error" @click="handleInitDefaultPermissions">
          Init
        </UButton>
        <UButton type="primary" @click="handleShowAddModule">
          Add Module
        </UButton>
      </div>
    </div>
    <div class="w-full text-center">
      <UTable
        :data="list"
        :loading="loading"
        :actions="bulkActions"
        tree
        associate-parent="uncheck"
        :default-expand-all="true"
      >
        <template #headers>
          <th scope="col" class="table-th">
            Name
          </th>
          <th scope="col" class="table-th">
            Uri
          </th>
          <th scope="col" class="table-th">
            Method
          </th>
          <th scope="col" class="table-th">
            Anonymous
          </th>
          <th scope="col" class="table-th">
            CreatedAt
          </th>
          <th scope="col" class="table-th" />
        </template>
        <template #rows="{ row }">
          <th colspan="5" scope="colgroup" class="whitespace-nowrap px-3.5 py-2 pl-4 text-left text-gray-500 sm:pl-6">
            {{ row.name }}
          </th>
          <th class="pl-4">
            <div class="flex items-center gap-4 text-xl">
              <button
                class="i-mdi:plus-circle text-indigo-500"
                @click="() => {
                  moduleForm.name = row.name;
                  initForm.parentId = row.id;
                  handleShowAdd();
                }"
              />
              <button
                class="i-mdi:book-edit text-green-500"
                @click="handleShowEditModule(row)"
              />
              <button
                class="i-mdi:delete text-red-500"
                @click="handleDelete(row.id)"
              />
            </div>
          </th>
        </template>
        <template #subs="{ sub }">
          <td class="table-td">
            {{ sub.name }}
          </td>
          <td class="table-td">
            {{ sub.uri }}
          </td>
          <td class="table-td">
            <UTag v-if="sub.parentId" :type="methodTypes[sub.method]">
              {{ sub.method }}
            </UTag>
          </td>
          <td class="table-td">
            <USwitch
              v-if="sub.parentId"
              :model-value="sub.anonymous"
              :loading="sub.toggleLoading"
              @update:model-value="handleToggleAnonymous(sub, $event)"
            />
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="sub.createdAt" />
          </td>
          <td class="table-td">
            <div class="flex items-center gap-4 text-xl">
              <button
                class="i-mdi:book-edit text-green-500"
                @click="handleShowEdit(sub)"
              />
              <button
                class="i-mdi:delete text-red-500"
                @click="handleDelete(sub.id)"
              />
            </div>
          </td>
        </template>
      </UTable>
    </div>
  </PageWrap>

  <UModal v-model="modalVisible">
    <template #header>
      Module:
      <span class="text-indigo-500">
        {{ moduleForm.name }}
      </span>
    </template>
    <div class="space-y-3">
      <div class="space-y-3">
        <UInput v-model="form.name" label="Name" />
        <UInput v-model="form.uri" label="Uri" />
        <p>
          Method
        </p>
        <div class="flex gap-3">
          <URadio
            v-for="item in methods" :key="item"
            v-model="form.method"
            :type="methodTypes[item]"
            :value="item"
          >
            {{ item }}
          </URadio>
        </div>
      </div>
    </div>
    <template #footer>
      <UButton v-if="form.id" type="warning" @click="handleEdit(form)">
        Edit
      </UButton>
      <UButton v-else type="primary" @click="handleAdd">
        Add
      </UButton>
      <UButton @click="modalVisible = false">
        Cancel
      </UButton>
    </template>
  </UModal>

  <UModal v-model="moduleModalVisible">
    <template #header>
      {{ moduleForm.id ? 'Edit Module' : 'Add Module' }}
    </template>
    <div class="space-y-3">
      <template v-if="!moduleForm.id">
        <div class="flex items-center gap-4">
          <UInput v-model="moduleForm.objectName" placeholder="WebObject Name" />
          <UButton type="secondary" @click="handleCreateWebobjectPermissions">
            Generate
          </UButton>
        </div>
        <UDivider dashed>
          OR
        </UDivider>
      </template>
      <div class="flex items-center gap-4">
        <UInput v-model="moduleForm.name" label="Module Name" />
      </div>
    </div>
    <template #footer>
      <UButton v-if="moduleForm.id" type="warning" @click="handleEditModule(moduleForm)">
        Edit
      </UButton>
      <UButton v-else type="primary" @click="handleAddModule">
        Add
      </UButton>
      <UButton @click="moduleModalVisible = false">
        Cancel
      </UButton>
    </template>
  </UModal>
</template>
