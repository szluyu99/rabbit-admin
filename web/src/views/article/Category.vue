<script setup>
import { onMounted } from 'vue'
import { UButton, UInput, UModal, UTable } from 'unocss-ui'

import Date from '@/components/Date.vue'
import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'

const {
  list, keyword, loading, modalVisible, form, page, limit, total,
  handleQuery, handleSearch,
  handleShowAdd, handleAdd,
  handleShowEdit, handleEdit,
  handleDelete, handleBatchDelete,
} = useTable({
  queryFn: params => request.post('/api/category', params),
  addFn: item => request.put('/api/category', item),
  deleteFn: id => request.delete(`/api/category/${id}`),
  editFn: item => request.patch(`/api/category/${item.id}`, item),
})

const bulkActions = [
  {
    name: 'Delete All',
    clicked: keys => handleBatchDelete(keys),
  },
]

onMounted(() => {
  handleQuery()
})
</script>

<template>
  <PageWrap title="Category Page">
    <div class="flex items-center justify-between">
      <div class="w-52">
        <UInput
          v-model="keyword"
          placeholder="Search name"
          right-icon="i-mdi:magnify"
          @click-right="handleSearch"
          @keyup.enter="handleSearch"
        />
      </div>
      <UButton type="primary" @click="handleShowAdd">
        Add Category
      </UButton>
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" :actions="bulkActions" header-color>
        <template #headers>
          <th scope="col" class="table-th">
            Name
          </th>
          <th scope="col" class="table-th">
            Article Count
          </th>
          <th scope="col" class="table-th">
            CreatedAt
          </th>
          <th scope="col" class="table-th">
            UpdatedAt
          </th>
          <th scope="col" class="table-th" />
        </template>
        <template #rows="{ row }">
          <td class="table-td font-bold">
            {{ row.name }}
          </td>
          <td class="table-td">
            {{ row.article_count }}
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="row.created_at" />
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="row.updated_at" />
          </td>
          <td class="table-td">
            <div class="flex items-center gap-4 text-xl">
              <span
                class="i-mdi:delete cursor-pointer text-red-400 hover:text-red-500"
                @click="handleDelete(row.id)"
              />
              <span
                class="i-mdi:book-edit cursor-pointer text-green-400 hover:text-green-500"
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
      {{ form.id ? 'Edit Category' : 'Add Category' }}
    </template>
    <div class="space-y-3">
      <UInput v-model="form.name" label="Name" />
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
</template>
