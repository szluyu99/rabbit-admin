// 更通用的 CRUD 操作封装
import { ref } from 'vue'
import { useConfirm } from './useConfirm'

interface CRUDOptions {
  initForm: {}
  validateForm?: () => boolean
  queryFn?: Function
  editFn?: Function
  addFn?: Function
  deleteFn?: Function
}

export default function useCRUD({
  initForm = {},
  addFn,
  editFn,
  deleteFn,
  queryFn,
}: CRUDOptions) {
  // table
  const selectedIds = ref<number[] | string[]>([])
  const loading = ref(false)
  const form = ref<any>({ ...initForm })
  const modalVisible = ref(false)
  const modalLoading = ref(false)

  const data = ref<any>(null)

  function handleShowAdd(): void {
    modalVisible.value = true
    form.value = { ...initForm }
  }

  function handleShowEdit(row: any): void {
    modalVisible.value = true
    form.value = { ...initForm, ...row }
  }

  async function handleQuery(): Promise<void> {
    if (!queryFn)
      return

    selectedIds.value = []
    loading.value = true
    try {
      const resp = await queryFn()
      data.value = resp
    }
    catch (err: any) {
      console.error(err)
    }
    finally {
      setTimeout(() => loading.value = false, 100)
    }
  }

  async function handleAdd(): Promise<void> {
    if (!addFn)
      return

    try {
      modalLoading.value = true
      await addFn()
      modalVisible.value = false
      setTimeout(() => handleQuery(), 100)
    }
    catch (err: any) {
      console.error(err)
    }
    finally {
      modalLoading.value = false
    }
  }

  async function handleEdit(item: any): Promise<void> {
    if (!editFn)
      return

    try {
      modalLoading.value = true
      await editFn(item)
      modalVisible.value = false
      setTimeout(() => handleQuery(), 100)
    }
    catch (err: any) {
      console.error(err)
    }
    finally {
      modalLoading.value = false
    }
  }

  async function handleDelete(id: number | string): Promise<void> {
    if (!deleteFn)
      return

    useConfirm({
      title: 'Delete Confirm',
      content: 'Are you sure to delete this item?',
      onOk: async () => {
        try {
          await deleteFn(id)
          setTimeout(() => handleQuery(), 100)
        }
        catch (err: any) {
          console.error(err)
        }
      },
    })
  }

  return {
    selectedIds,
    loading,
    form,
    modalVisible,
    modalLoading,
    data,
    handleShowAdd,
    handleShowEdit,
    handleQuery,
    handleAdd,
    handleEdit,
    handleDelete,
  }
}
