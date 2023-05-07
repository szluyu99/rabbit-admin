<script setup>
import { onMounted } from 'vue'
import Table from '@/components/base/Table.vue'
import Input from '@/components/base/Input.vue'
import Pagination from '@/components/Pagination.vue'
import PageCard from '@/components/PageCard.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import { formatDate } from '@/utils/helper'

const {
  list, keyword, loading, pos, limit, total,
  handleQuery, handleSearch,
} = useTable({
  queryFn: params => request.post('/api/user', params),
  addFn: item => request.put('/api/user', item),
  deleteFn: id => request.delete(`/api/user/${id}`),
  editFn: item => request.patch(`/api/user/${item.id}`, item),
  batchFn: ids => request.delete('/api/user', ids),
})

onMounted(() => {
  handleQuery()
})
</script>

<template>
  <PageCard title="User Page">
    <div class="flex justify-between items-center">
      <div class="w-64 sm:w-xs">
        <Input
          v-model="keyword"
          suff-icon="i-mdi:magnify"
          @append="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
    </div>
    <div class="w-full text-center">
      <Table :data="list" :loading="loading">
        <template #headers>
          <th scope="col" class="py-3.5 px-3 col-span-3 text-sm font-semibold text-gray-900 sm:pl-6">
            Email
          </th>
          <th scope="col" class="py-3.5 px-3 col-span-3 text-sm font-semibold text-gray-900 sm:pl-6">
            Enabled
          </th>
          <th scope="col" class="py-3.5 px-3 col-span-3 text-sm font-semibold text-gray-900 sm:pl-6">
            LastLogin
          </th>
          <th scope="col" class="py-3.5 px-3 col-span-3 text-sm font-semibold text-gray-900 sm:pl-6">
            LastLoginIP
          </th>
          <th scope="col" class="px-4 py-3.5 col-span-3 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 sm:pl-6 text-sm text-gray-500">
            {{ row.email }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 sm:pl-6 text-sm text-gray-500">
            {{ row.enabled }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 sm:pl-6 text-sm text-gray-500">
            {{ formatDate(row.lastLogin) }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 sm:pl-6 text-sm text-gray-500">
            {{ row.lastLoginIP }}
          </td>
          <td class="w-32 whitespace-nowrap px-3.5 py-2 text-lg text-gray-500 group">
            <div class="flex items-center gap-2 lg:px-4">
              <span class="i-mdi:human-edit cursor-pointer text-red-500" />
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
</template>
