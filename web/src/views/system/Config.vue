<script setup>
import { onMounted } from 'vue'
import { UButton, UInput, UModal, UTable } from 'unocss-ui'

import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'

const {
  list, keyword, loading, modalVisible, form, page, limit, total,
  handleQuery, handleSearch, handleShowEdit, handleEdit, handleDelete,
} = useTable({
  queryFn: params => request.post('/api/config', params),
  addFn: item => request.put('/api/config', item),
  deleteFn: id => request.delete(`/api/config/${id}`),
  editFn: item => request.patch(`/api/config/${item.id}`, item),
})

onMounted(() => {
  handleQuery()
})
</script>

<template>
  <PageWrap title="System Config">
    <template #description>
      key-value pair configuration stored in the database.
    </template>
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
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" header-color>
        <template #headers>
          <th scope="col" class="table-th">
            Key
          </th>
          <th scope="col" class="table-th">
            Value
          </th>
          <th scope="col" class="table-th">
            Desc
          </th>
          <th scope="col" class="col-span-3 px-4 py-3.5 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="table-td">
            {{ row.key }}
          </td>
          <td class="table-td">
            {{ row.value }}
          </td>
          <td class="table-td">
            {{ row.desc }}
          </td>
          <td class="table-td">
            <div class="flex items-center gap-4 text-xl">
              <span
                class="i-mdi:delete cursor-pointer text-red-500"
                @click="handleDelete(row.id)"
              />
              <span
                class="i-mdi:book-edit cursor-pointer text-green-500"
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
      {{ form.id ? 'Edit Config' : 'Add Config' }}
    </template>
    <div class="space-y-3">
      <UInput v-model="form.key" label="Key" />
      <UInput v-model="form.value" label="Value" />
      <UInput v-model="form.desc" label="Desc" />
    </div>
    <template #footer>
      <UButton v-if="form.id" type="warning" @click="handleEdit(form)">
        Edit
      </UButton>
      <UButton @click="modalVisible = false">
        Cancel
      </UButton>
    </template>
  </UModal>
</template>
