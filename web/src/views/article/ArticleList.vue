<script setup>
import { computed, onMounted, ref } from 'vue'
import { UButton, UInput, USelect, USwitch, UTable, UTag } from 'unocss-ui'

import Date from '@/components/Date.vue'
import Pagination from '@/components/Pagination.vue'
import PageWrap from '@/components/PageWrap.vue'
import { useConfirm } from '@/composables/useConfirm'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import api from '@/api'

const tabs = ref([
  { name: 'All', value: '', current: true },
  { name: 'Public', value: 'public', current: false },
  { name: 'Private', value: 'private', current: false },
  { name: 'Draft', value: 'draft', current: false },
  { name: 'Trash', value: 'trash', current: false },
])

const currentTab = computed(() => tabs.value.find(e => e.current).value)

const {
  list, keyword, loading, page, limit, total, queryParams,
  handleQuery, handleSearch, handleDelete, handleBatchDelete,
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
    is_delete: false,
  },
})

onMounted(() => {
  handleQuery()
})

const bulkActions = [
  {
    name: 'Delete All',
    clicked: (keys) => {
      currentTab.value === 'trash' ? handleBatchDelete(keys) : handleBatchSoftDelete(keys)
    },
  },
  {
    name: 'Export All',
    clicked: () => {},
  },
]

async function handleToggleTop(row, val) {
  try {
    row.toggleLoading = true
    await request.patch(`/api/article/${row.id}`, { is_top: val })
    setTimeout(() => {
      row.toggleLoading = false
      row.is_top = val
    }, 300)
  }
  catch (err) {
    row.toggleLoading = false
  }
}

function handleBatchSoftDelete(ids) {
  useConfirm({
    title: 'Delete Confirm',
    content: 'Are you sure to delete the selected articles? They will be moved to trash.',
    onOk: async () => {
      await Promise.all(ids.map(id => request.patch(`/api/article/${id}`, { is_delete: true })))
      handleQuery()
    },
  })
}

function handleSoftDelete(row) {
  useConfirm({
    title: 'Delete Confirm',
    content: `Are you sure to delete the article: <span class="font-bold">${row.title}</span>? It will be moved to trash.`,
    onOk: async () => {
      await request.patch(`/api/article/${row.id}`, { is_delete: true })
      handleQuery()
    },
  })
}

async function handleCancelSoftDelete(row) {
  await request.patch(`/api/article/${row.id}`, { is_delete: false })
  handleQuery()
}

const tagOptions = ref([])
const categoryOptions = ref([])
const typeOptions = [
  { label: 'original', value: 'original' },
  { label: 'reprint', value: 'reprint' },
  { label: 'translation', value: 'translation' },
]

onMounted(async () => {
  tagOptions.value = await api.getTagOptions()
  categoryOptions.value = await api.getCategoryOptions()
})

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

const typeTypes = {
  original: 'primary',
  reprint: 'secondary',
  translation: 'accent',
}

const statusTypes = {
  public: 'success',
  private: 'warning',
}
</script>

<template>
  <PageWrap title="Article Page">
    <div>
      <div class="sm:hidden">
        <label for="tabs" class="sr-only">USelect a tab</label>
        <!-- Use an "onChange" listener to redirect the user to the USelected tab URL. -->
        <select
          id="tabs"
          name="tabs"
          class="block w-full border-gray-300 rounded-md py-1 pl-3 pr-10 text-base focus:border-indigo-500 sm:text-sm focus:outline-none focus:ring-indigo-500"
        >
          <option
            v-for="tab in tabs"
            :key="tab.name"
            :selected="tab.current"
            @change="handleClickTab($event)"
          >
            {{ tab.name }}
          </option>
        </select>
      </div>
      <div class="hidden sm:block">
        <div class="border-b border-gray-200">
          <nav class="flex -mb-px space-x-5" aria-label="Tabs">
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
      <div class="w-full flex flex-wrap items-center gap-1 sm:gap-3">
        <div class="w-52">
          <UInput
            v-model="keyword"
            placeholder="Search title"
            right-icon="i-mdi:magnify"
            @click-right="handleSearch"
            @keyup.enter="handleSearch"
          />
        </div>
        <div class="w-36">
          <USelect
            v-model="queryParams.type"
            :options="typeOptions"
            placeholder="Type"
            clearable
            @update:model-value="handleQuery"
          />
        </div>
        <div class="w-36">
          <USelect
            v-model="queryParams.tag_id"
            :options="tagOptions"
            placeholder="Tag"
            clearable
            @update:model-value="handleQuery"
          />
        </div>
        <div class="w-36">
          <USelect
            v-model="queryParams.category_id"
            :options="categoryOptions"
            placeholder="Category"
            clearable
            @update:model-value="handleQuery"
          />
        </div>
      </div>
      <UButton type="primary" @click="$router.push('/article/write')">
        Write Article
      </UButton>
    </div>
    <div class="w-full text-center">
      <UTable :data="list" :loading="loading" :actions="bulkActions" header-color>
        <template #headers>
          <th scope="col" class="table-th">
            Cover
          </th>
          <th scope="col" class="table-th max-w-64">
            Title
          </th>
          <th scope="col" class="table-th">
            Category
          </th>
          <th scope="col" class="table-th">
            Tags
          </th>
          <th scope="col" class="table-th">
            Type
          </th>
          <th scope="col" class="table-th">
            UpdatedAt
          </th>
          <th scope="col" class="table-th">
            IsTop
          </th>
          <th scope="col" class="table-th">
            Status
          </th>
          <th scope="col" class="col-span-3 px-4 py-3.5 text-sm font-semibold text-gray-900" />
        </template>
        <template #rows="{ row }">
          <td class="table-td">
            <img class="max-w-40 min-w-20" :src="row.cover" alt="">
          </td>
          <td class="table-td max-w-64">
            {{ row.title }}
          </td>
          <td class="table-td">
            {{ row.category?.name }}
          </td>
          <td class="table-td">
            <div class="flex gap-2">
              <UTag v-for="tag of row.tags" :key="tag.id" type="success">
                {{ tag.name }}
              </UTag>
            </div>
          </td>
          <td class="table-td">
            <UTag :type="typeTypes[row.type]">
              {{ row.type }}
            </UTag>
          </td>
          <td class="table-td whitespace-nowrap">
            <Date :value="row.updated_at" />
          </td>
          <td class="table-td">
            <div class="flex items-center">
              <USwitch
                :model-value="row.is_top"
                :loading="row.toggleLoading"
                @update:model-value="handleToggleTop(row, $event)"
              />
            </div>
          </td>
          <td class="table-td">
            <UTag :type="statusTypes[row.status]">
              {{ row.status }}
            </UTag>
          </td>
          <td class="table-td w-32">
            <div class="flex items-center gap-4 text-xl">
              <span
                v-if="!row.is_delete"
                class="i-mdi:delete cursor-pointer text-red-400 hover:text-red-500"
                @click="handleSoftDelete(row)"
              />
              <span
                v-else
                class="i-mdi:backburger cursor-pointer text-green-400 hover:text-green-500"
                @click="handleCancelSoftDelete(row)"
              />
              <span
                class="i-mdi:pen cursor-pointer text-primary-400 hover:text-primary-500"
                @click="$router.push(`/article/write/${row.id}`)"
              />
              <span
                v-if="row.is_delete"
                class="i-mdi:delete cursor-pointer text-red-400 hover:text-red-500"
                @click="handleDelete(row.id)"
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
</template>
