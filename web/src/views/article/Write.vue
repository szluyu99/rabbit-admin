<script setup>
import { computed, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { marked } from 'marked'
import { debounce } from 'lodash-es'

import Input from '@/components/base/Input.vue'
import Button from '@/components/base/Button.vue'
import Select from '@/components/base/Select.vue'
import Switch from '@/components/base/Switch.vue'
import Radio from '@/components/base/Radio.vue'
import Modal from '@/components/base/Modal.vue'
import DynamicTags from '@/components/base/DynamicTags.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'
import api from '@/api'

const router = useRouter()
const route = useRoute()

const {
  modalVisible, form,
  handleShowAdd, handleShowEdit,
} = useTable({
  initForm: {
    title: '',
    content: '',
    category_id: 0,
    tag_id: 0,
    type: '',
    status: '',
    is_top: false,
    tag_names: [],
  },
})

const articleId = ref('')
const title = ref('')
const content = ref('# Hello')
const output = computed(() => marked(content.value))
const updateContent = debounce(e => content.value = e.target.value, 100)

function handlePublish() {
  articleId.value ? handleShowEdit(form.value) : handleShowAdd()
}

async function handleSaveOrUpdate() {
  form.value.title = title.value
  form.value.content = content.value

  try {
    await request.post('/api/article/saveOrUpdate', form.value)
    router.push('/article/list')
  }
  catch (e) {
    console.error(e)
  }
}

// const tagOptions = ref([])
const categoryOptions = ref([])

onMounted(() => {
  articleId.value = route.params.articleId

  if (articleId.value) {
    request.get(`/api/article/${articleId.value}`).then((res) => {
      title.value = res.title
      content.value = res.content
      form.value = res
      form.value.tag_names = res.tags.map(e => e.name)
    })
  }

  api.categoryOptions().then((res) => {
    categoryOptions.value = (res.items || []).map(e => ({
      label: e.name,
      value: e.id,
    }))
  })
  // api.tagOptions().then((res) => {
  //   tagOptions.value = (res.items || []).map(e => ({
  //     label: e.name,
  //     value: e.id,
  //   }))
  // })
})

const typeOptions = [
  { label: 'original', value: 'original' },
  { label: 'reprint', value: 'reprint' },
  { label: 'translation', value: 'translation' },
]
</script>

<template>
  <div class="max-w-7xl rounded bg-white p-5 shadow space-y-4">
    <div class="flex items-center justify-between gap-4">
      <Input v-model="title" placeholder="Article title" />
      <Button type="warning">
        Save Draft
      </Button>
      <Button type="success" @click="handlePublish">
        Publish
      </Button>
    </div>
    <div class="h-4xl flex overflow-y-auto">
      <textarea
        class="box-border w-full resize-none overflow-scroll border-0 border-r-1 border-solid bg-gray-50 p-2 outline-none sm:w-1/2"
        :value="content"
        @input="updateContent"
      />
      <article class="box-border hidden w-1/2 overflow-scroll p-2 px-5 prose prose-truegray sm:block xl:text-xl">
        <div class="text-base" v-html="output" />
      </article>
    </div>
  </div>
  <Modal v-model="modalVisible">
    <p class="mb-5 text-xl font-bold">
      Publish article
    </p>
    <Select
      v-model="form.category_id"
      label="Category"
      :options="categoryOptions"
    />
    <div class="my-3 space-y-2">
      <div class="text-sm font-medium leading-6 text-gray-900">
        Article Tags
      </div>
      <div class="h-5">
        <DynamicTags v-model="form.tag_names" :max="3" />
      </div>
    </div>
    <Select
      v-model="form.type"
      label="Article Type"
      :options="typeOptions"
    />
    <div class="space-y-2">
      <span class="text-sm font-medium leading-6 text-gray-900">
        Article Cover
      </span>
      <Input v-model="form.cover" />
    </div>
    <div class="my-3 flex gap-4">
      <Switch v-model="form.is_top">
        Is Top
      </Switch>
    </div>
    <div class="flex items-center gap-4 space-x-2">
      <span class="text-sm font-medium leading-6 text-gray-900">
        Article Status
      </span>
      <Radio v-model="form.status" value="public">
        Public
      </Radio>
      <Radio v-model="form.status" value="private">
        Private
      </Radio>
    </div>
    <div class="flex flex-row justify-end gap-2">
      <Button @click="modalVisible = false">
        Cancel
      </Button>
      <Button type="primary" @click="handleSaveOrUpdate">
        Save
      </Button>
    </div>
  </Modal>
</template>
