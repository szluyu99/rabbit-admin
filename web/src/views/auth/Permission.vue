<script setup>
import { onMounted } from 'vue'

import Table from '@/components/base/Table.vue'
import Input from '@/components/base/Input.vue'
import Button from '@/components/base/Button.vue'
import Modal from '@/components/base/Modal.vue'
import Pagination from '@/components/Pagination.vue'
import PageCard from '@/components/PageCard.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import { formatDate } from '@/utils/helper'

const {
  list, keyword, loading, modalVisible, form, pos, limit, total,
  handleQuery, handleSearch,
  handleShowAdd, handleAdd,
  handleShowEdit, handleEdit,
  handleDelete, handleBatch,
} = useTable({
  queryFn: params => request.post('/api/permission', params),
  addFn: item => request.put('/api/permission', item),
  deleteFn: id => request.delete(`/api/permission/${id}`),
  editFn: item => request.patch(`/api/permission/${item.id}`, item),
  batchFn: ids => request.delete('/api/permission', ids),
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
  <PageCard title="Permission Page">
    <div class="flex items-center justify-between">
      <div class="w-64 sm:w-xs">
        <Input
          v-model="keyword"
          suff-icon="i-mdi:magnify"
          @append="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
      <div>
        <Button type="primary" @click="handleShowAdd">
          Add
        </Button>
      </div>
    </div>
    <div class="w-full text-center">
      <Table :data="list" :loading="loading" :actions="bulkActions">
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
            Code
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            P1
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            P2
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            P3
          </th>
          <th scope="col" class="col-span-3 px-4 py-3.5 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="w-16 whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ formatDate(row.createdAt) }}
          </td>
          <td class="w-16 whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ formatDate(row.updatedAt) }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.name }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.code }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.p1 }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.p2 }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.p3 }}
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
      </Table>
    </div>
    <Pagination
      v-model:pos="pos"
      v-model:limit="limit"
      :total="total"
      @query="handleQuery"
    />
  </PageCard>
  <Modal v-model="modalVisible">
    <div class="mb-5 text-lg font-bold">
      {{ form.id ? 'Edit Permission' : 'Add Permission' }}
    </div>
    <div class="space-y-3">
      <Input v-model="form.name" label="Name" />
      <Input v-model="form.code" label="Code" />
      <Input v-model="form.p1" label="P1" />
      <Input v-model="form.p2" label="P2" />
      <Input v-model="form.p3" label="P3" />
    </div>
    <div class="mt-5 sm:mt-4 sm:flex sm:flex-row-reverse sm:gap-3">
      <Button v-if="form.id" type="success" @click="handleEdit(form)">
        Edit
      </Button>
      <Button v-else type="primary" @click="handleAdd">
        Add
      </Button>
      <Button class="mt-3 sm:mt-0" @click="modalVisible = false">
        Cancel
      </Button>
    </div>
  </Modal>
</template>
