<script setup>
import { onMounted } from 'vue'
import { UButton, UInput, UModal, UTable } from 'unocss-ui'

import Date from '@/components/Date.vue'
import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'

const {
  list, keyword, loading, modalVisible, form, pos, limit, total,
  handleQuery, handleSearch,
  handleShowAdd, handleAdd,
  handleShowEdit, handleEdit,
  handleDelete, handleBatch,
} = useTable({
  queryFn: params => request.post('/api/group', params),
  addFn: item => request.put('/api/group', item),
  deleteFn: id => request.delete(`/api/group/${id}`),
  editFn: item => request.patch(`/api/group/${item.id}`, item),
  batchFn: ids => request.delete('/api/group', ids),
})

const bulkActions = [
  {
    name: 'Delete All',
    clicked: keys => handleBatch(keys),
  },
]

onMounted(() => {
  handleQuery()
})
</script>

<template>
  <PageWrap title="Group Page">
    <div class="flex items-center justify-between">
      <div class="w-64 sm:w-xs">
        <UInput
          v-model="keyword"
          suff-icon="i-mdi:magnify"
          @append="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
      <div>
        <UButton type="primary" @click="handleShowAdd">
          Add
        </UButton>
      </div>
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" :actions="bulkActions">
        <template #headers>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            CreatedAt
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            UpdatedAt
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Name
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Extra
          </th>
          <th scope="col" class="col-span-3 px-4 py-3.5 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="w-16 whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <Date :value="row.created_at" />
          </td>
          <td class="w-16 whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <Date :value="row.updated_at" />
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.name }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.extra }}
          </td>
          <td class="group w-32 whitespace-nowrap px-3.5 py-2 text-lg text-gray-500">
            <div class="flex items-center gap-2 lg:px-4">
              <span
                class="i-mdi:delete invisible cursor-pointer text-red-500 group-hover:visible"
                @click="handleDelete(row.id)"
              />
              <span
                class="i-mdi:book-edit invisible cursor-pointer text-green-500 group-hover:visible"
                @click="handleShowEdit(row)"
              />
            </div>
          </td>
        </template>
      </UTable>
    </div>
    <Pagination
      v-model:pos="pos"
      v-model:limit="limit"
      :total="total"
      @query="handleQuery"
    />
  </PageWrap>

  <UModal v-model="modalVisible">
    <div class="mb-5 text-lg font-bold">
      {{ form.id ? 'Edit Group' : 'Add Group' }}
    </div>
    <div class="space-y-3">
      <UInput v-model="form.name" label="Name" />
    </div>
    <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse sm:gap-3">
      <UButton v-if="form.id" type="success" @click="handleEdit(form)">
        Edit
      </UButton>
      <UButton v-else type="primary" @click="handleAdd">
        Add
      </UButton>
      <UButton class="mt-3 sm:mt-0" @click="modalVisible = false">
        Cancel
      </UButton>
    </div>
  </UModal>
</template>
