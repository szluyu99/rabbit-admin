import { reactive, ref } from 'vue'
import { useConfirm } from './useConfirm'

type FilterOp = '=' | '<>' | 'in' | 'not_in' | '>' | '>=' | '<' | '<='
type OrderOp = 'asc' | 'desc'

interface Filter {
  name: string
  op: FilterOp
  value: any
}

interface Order {
  name: string
  op: OrderOp
}

interface ListOptions {
  initForm: {}
  validateForm?: () => boolean
  queryFn?: ((params: QueryParams) => Promise<QueryResult>) | Function | null
  editFn?: (item: EditItem) => Promise<void>
  addFn?: (form: any) => Promise<void> // | boolean
  deleteFn?: (id: string | number) => Promise<void>
  batchFn?: (form: string[]) => Promise<void>
  afterQuery: Function | null
  extraParams: any
}

interface EditItem {
  id: string | number
  // ...
}

export interface QueryParams {
  page?: number
  limit?: number
  keyword?: string
  filters?: Array<Filter>
  orders?: Array<Order>
}

interface QueryResult {
  page?: number
  limit?: number
  keyword?: string
  total?: number
  items?: any[]
}

export default function useTable({
  initForm = {},
  addFn,
  editFn,
  deleteFn,
  batchFn,
  queryFn,
  afterQuery,
  extraParams = {},
}: ListOptions) {
  // table
  const selectedIds = ref<number[] | string[]>([])
  const loading = ref(false)
  const form = ref<any>({ ...initForm })
  const modalVisible = ref(false)
  const modalLoading = ref(false)

  // queryParams
  const page = ref(1)
  const limit = ref(10)
  const keyword = ref('')
  const queryParams = reactive(extraParams)

  const filters = ref<Filter[]>([])
  const orders = ref<Order[]>([])

  // queryResult
  const total = ref(0)
  const list = ref<any[]>([])

  function handleShowAdd(): void {
    modalVisible.value = true
    form.value = { ...initForm }
  }

  function handleShowEdit(row: any): void {
    modalVisible.value = true
    form.value = { ...initForm, ...row }
  }

  async function handleQuery(): Promise<void> {
    // console.log(typeof queryFn)
    if (!queryFn)
      return

    selectedIds.value = []
    loading.value = true
    try {
      const resp = await queryFn({
        page: page.value,
        limit: limit.value,
        keyword: keyword.value,
        filters: filters.value,
        orders: orders.value,
        ...queryParams,
      })

      if (resp?.total)
        total.value = resp.total
      if (resp?.items)
        list.value = resp.items ?? []

      afterQuery && afterQuery()

      // Empty data found and not on the first page
      if (list.value.length === 0 && page.value !== 1) {
        page.value -= 1
        await handleQuery()
      }
    }
    catch (err: any) {
      console.error(err)
    }
    finally {
      setTimeout(() => loading.value = false, 100)
    }
  }

  function handleSearch(): void {
    page.value = 1
    handleQuery()
  }

  async function handleEdit(item: EditItem): Promise<void> {
    if (!editFn)
      return

    try {
      modalLoading.value = true
      await editFn(item)
      modalVisible.value = false
      handleQuery()
    }
    catch (err: any) {
      console.error(err)
    }
    finally {
      modalLoading.value = false
    }
  }

  async function handleAdd(): Promise<void> {
    if (!addFn)
      return

    try {
      modalLoading.value = true
      await addFn(form.value)
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
      content: 'Are you sure you want to delete this? This action cannot be undone.',
      onOk: async () => {
        try {
          await deleteFn(id)
          modalVisible.value = false
          handleQuery()
        }
        catch (err: any) {
          console.error(err)
        }
      },
    })
  }

  // TODO: optimize handleBatch && handleBatchDelete
  async function handleBatch(ids: number[] | string[]): Promise<void> {
    if (!batchFn)
      return

    useConfirm({
      title: 'Delete Confirm',
      content: 'Are you sure you want to delete all? This action cannot be undone.',
      onOk: async () => {
        try {
          await batchFn(ids.map(e => String(e)))
          handleQuery()
        }
        catch (err: any) {
          console.error(err)
        }
      },
    })
  }

  async function handleBatchDelete(ids: number[] | string[]): Promise<void> {
    if (!deleteFn)
      return

    useConfirm({
      title: 'Delete Confirm',
      content: 'Are you sure you want to delete all? This action cannot be undone.',
      onOk: async () => {
        try {
          await Promise.all(ids.map(id => deleteFn(id)))
        }
        catch (err) {
          console.error(err)
        }
        finally {
          handleQuery()
        }
      },
    })
  }

  function handleOrder(field: string, op: OrderOp) {
    orders.value = orders.value.filter(e => e.name !== field)
    orders.value.push({ name: field, op })
    handleQuery()
  }

  function handleRemoveOrder(field: string) {
    orders.value = orders.value.filter(e => e.name !== field)
    handleQuery()
  }

  function handleFilter(field: string, op: FilterOp, value: any) {
    filters.value = filters.value.filter(e => e.name !== field)
    filters.value.push({ name: field, op, value })
    handleQuery()
  }

  function handleRemoveFilter(field: string) {
    filters.value = filters.value.filter(e => e.name !== field)
    handleQuery()
  }

  function handleReset() {
    page.value = 1
    limit.value = 10
    keyword.value = ''
    filters.value = []
    orders.value = []

    handleQuery()
  }

  return {
    page,
    limit,
    keyword,
    filters,
    orders,
    total,
    list,
    loading,
    selectedIds,
    modalVisible,
    modalLoading,
    form,
    queryParams,
    handleQuery,
    handleSearch,
    handleOrder,
    handleRemoveOrder,
    handleFilter,
    handleRemoveFilter,
    handleEdit,
    handleShowEdit,
    handleAdd,
    handleShowAdd,
    handleDelete,
    handleBatch,
    handleBatchDelete,
    handleReset,
  }
}
