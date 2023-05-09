<script setup>
import { onMounted, ref } from 'vue'

import Table from '@/components/base/Table.vue'
import Input from '@/components/base/Input.vue'
import Select from '@/components/base/Select.vue'
import Button from '@/components/base/Button.vue'
import Pagination from '@/components/Pagination.vue'
import PageCard from '@/components/PageCard.vue'
import Tag from '@/components/base/Tag.vue'
import Switch from '@/components/base/Switch.vue'

import { formatDate } from '@/utils/helper'
import useTable from '@/composables/useTable'
import request from '@/api/request'
import api from '@/api'

const {
  list, keyword, loading, page, limit, total, queryParams,
  handleQuery, handleSearch,
  handleDelete, handleBatch,
} = useTable({
  queryFn: params => request.post('/api/article', params),
  addFn: item => request.put('/api/article', item),
  deleteFn: id => request.delete(`/api/article/${id}`),
  editFn: item => request.patch(`/api/article/${item.id}`, item),
  batchFn: ids => request.delete('/api/article', ids),
  extraParams: {
    tag_id: 0,
    category_id: 0,
    type: '',
  },
})

const bulkActions = [
  {
    name: 'Delete All',
    clicked: keys => handleBatch(keys),
  },
  {
    name: 'Export All',
    clicked: () => {},
  },
]

async function toggleIsStop(id, val) {
  try {
    await request.patch(`/api/article/${id}`, { is_top: val })
    handleQuery()
  }
  catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  handleQuery()
})

const tagOptions = ref([])
const categoryOptions = ref([])
const typeOptions = [
  { label: 'original', value: 'original' },
  { label: 'reprint', value: 'reprint' },
  { label: 'translation', value: 'translation' },
]

onMounted(() => {
  api.tagOptions().then((res) => {
    tagOptions.value = (res.items || []).map(e => ({
      label: e.name,
      value: e.id,
    }))
  })
  api.categoryOptions().then((res) => {
    categoryOptions.value = (res.items || []).map(e => ({
      label: e.name,
      value: e.id,
    }))
  })
})

function typeTag(type) {
  switch (type) {
    case 'original':
      return 'primary'
    case 'reprint':
      return 'secondary'
    case 'translation':
      return 'accent'
    default:
      return 'default'
  }
}

const tabs = ref([
  { name: 'All', value: '', current: true },
  { name: 'Public', value: 'public', current: false },
  { name: 'Private', value: 'private', current: false },
  { name: 'Draft', value: 'draft', current: false },
  { name: 'Trash', value: 'trash', current: false },
])

function handleClickTab(tab) {
  tabs.value.forEach(e => e.current = e.name === tab.name)
  if (tab.value === 'trash') {
    queryParams.is_delete = true
    queryParams.status = ''
  }
  else {
    queryParams.is_delete = false
    queryParams.status = tab.value
  }
  handleQuery()
}
</script>

<template>
  <PageCard title="Article Page">
    <div>
      <div class="sm:hidden">
        <label for="tabs" class="sr-only">Select a tab</label>
        <!-- Use an "onChange" listener to redirect the user to the selected tab URL. -->
        <select
          id="tabs"
          name="tabs"
          class="block w-full border-gray-300 rounded-md py-1.5 pl-3 pr-10 text-base focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500"
        >
          <option
            v-for="tab in tabs" :key="tab.name" :selected="tab.current"
            @change="handleClickTab($event)"
          >
            {{ tab.name }}
          </option>
        </select>
      </div>
      <div class="hidden sm:block">
        <div class="border-b border-gray-200">
          <nav class="flex -mb-px space-x-8" aria-label="Tabs">
            <span class="cursor-default whitespace-nowrap px-1 py-4 text-sm font-medium">
              Status
            </span>
            <span
              v-for="tab in tabs" :key="tab.name"
              class="cursor-pointer whitespace-nowrap border-b-2 px-1 py-4 text-sm font-medium"
              :class="[tab.current ? 'border-indigo-500 text-indigo-600' : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700']"
              :aria-current="tab.current ? 'page' : undefined"
              @click="handleClickTab(tab)"
            >
              {{ tab.name }}
            </span>
          </nav>
        </div>
      </div>
    </div>
    <div class="flex items-center justify-between">
      <div class="w-full flex flex-wrap items-center gap-1 sm:flex sm:gap-4">
        <div class="w-56">
          <Input
            v-model="keyword"
            suff-icon="i-mdi:magnify"
            placeholder="Search title"
            @append="handleSearch"
            @keyup.enter="handleSearch"
          />
        </div>
        <div class="w-44">
          <Select
            v-model="queryParams.type"
            :options="typeOptions"
            placeholder="Type"
            clearable
            @change="handleQuery"
          />
        </div>
        <div class="w-44">
          <Select
            v-model="queryParams.tag_id"
            :options="tagOptions"
            placeholder="Tag"
            clearable
            @change="handleQuery"
          />
        </div>
        <div class="w-44">
          <Select
            v-model="queryParams.category_id"
            :options="categoryOptions"
            placeholder="Category"
            clearable
            @change="handleQuery"
          />
        </div>
      </div>
      <div>
        <Button type="primary" @click="$router.push('/article/write')">
          Write Article
        </Button>
      </div>
    </div>
    <div class="w-full text-center">
      <Table :data="list" :loading="loading" :actions="bulkActions">
        <template #headers>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Cover
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Title
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Category
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Tags
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            Type
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            UpdatedAt
          </th>
          <th scope="col" class="col-span-3 px-3 py-3.5 text-sm font-semibold text-gray-900 sm:pl-6">
            IsTop
          </th>
          <th scope="col" class="col-span-3 px-4 py-3.5 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <img class="w-32" :src="row.cover" alt="">
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.title }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            {{ row.category?.name }}
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <Tag v-for="tag of row.tags" :key="tag.id" type="success" class="ml-2">
              {{ tag.name }}
            </Tag>
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <Tag :type="typeTag(row.type)">
              {{ row.type }}
            </Tag>
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <div class="flex items-center justify-center gap-1">
              <span class="i-mdi:calendar-clock text-lg text-gray-400" />
              {{ formatDate(row.updated_at) }}
            </div>
          </td>
          <td class="whitespace-nowrap px-3.5 py-2 pl-4 text-sm text-gray-500 sm:pl-6">
            <div class="flex items-center justify-center">
              <Switch v-model="row.is_top" @update:model-value="toggleIsStop(row.id, $event)" />
            </div>
          </td>
          <td class="w-32 whitespace-nowrap px-3.5 py-2 text-xl text-gray-500">
            <div class="flex items-center gap-5 lg:px-4">
              <span
                class="i-mdi:delete cursor-pointer text-red-400 hover:text-red-500"
                @click="handleDelete(row.id)"
              />
              <span
                class="i-mdi:pen cursor-pointer text-primary-400 hover:text-primary-500"
                @click="$router.push(`/article/write/${row.id}`)"
              />
            </div>
          </td>
        </template>
      </Table>
    </div>
    <Pagination
      v-model:page="page"
      v-model:limit="limit"
      :total="total"
      @query="handleQuery"
    />
  </PageCard>
</template>
