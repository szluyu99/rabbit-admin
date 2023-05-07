<script setup>
import { marked } from 'marked'
import { computed, ref } from 'vue'

import Input from '@/components/base/Input.vue'
import Button from '@/components/base/Button.vue'

import useTable from '@/composables/useTable'
import request from '@/api/request'

const render = new marked.Renderer()
marked.setOptions({
  renderer: render,
  gfm: true,
  breaks: true,
  pedantic: false,
  sanitize: false,
  smartLists: true,
  smartypants: false,
})

const {
  handleAdd,
} = useTable({
  addFn: item => request.put('/api/article', item),
})

const input = ref(`
Marked - Markdown Parser
========================

[Marked] lets you convert [Markdown] into HTML.  Markdown is a simple text format whose goal is to be very easy to read and write, even when not converted to HTML.  This demo page will let you type anything you like and see how it gets converted.  Live.  No more waiting around.

How To Use The Demo
-------------------

1. Type in stuff on the left.
2. See the live updates on the right.

That's it.  Pretty simple.  There's also a drop-down option above to switch between various views:

- **Preview:**  A live display of the generated HTML as it would render in a browser.
- **HTML Source:**  The generated HTML before your browser makes it pretty.
- **Lexer Data:**  What [marked] uses internally, in case you like gory stuff like this.
- **Quick Reference:**  A brief run-down of how to format things using markdown.

Why Markdown?
-------------

It's easy.  It's not overly bloated, unlike HTML.  Also, as the creator of [markdown] says,

> The overriding design goal for Markdown's
> formatting syntax is to make it as readable
> as possible. The idea is that a
> Markdown-formatted document should be
> publishable as-is, as plain text, without
> looking like it's been marked up with tags
> or formatting instructions.

Ready to start writing?  Either start changing stuff on the left or
[clear everything](/demo/?text=) with a simple click.

[Marked]: https://github.com/markedjs/marked/
[Markdown]: http://daringfireball.net/projects/markdown/`)
const output = computed(() => marked(input.value))

// const update = debounce((e) => {
//   input.value = e.target.value
// }, 100)
</script>

<template>
  <div class="max-w-6xl rounded bg-white p-5 shadow space-y-4">
    <div class="flex items-center justify-between gap-4">
      <Input placeholder="Article title" />
      <Button type="warning">
        Save Draft
      </Button>
      <Button type="success">
        Publish
      </Button>
    </div>
    <div class="h-4xl flex overflow-y-auto">
      <textarea
        class="box-border w-full resize-none overflow-scroll border-0 border-r-1 border-solid bg-gray-50 p-2 outline-none sm:w-1/2"
        :value="input"
        @input="input = $event.target.value"
      />
      <article class="prose prose-truegray box-border hidden w-1/2 overflow-scroll p-2 px-5 sm:block xl:text-xl">
        <div class="text-base" v-html="output" />
      </article>
    </div>
  </div>
</template>
