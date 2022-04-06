<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="端口">
          <el-input
            v-model.trim="searchInfo.port"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="host">
          <el-input
            v-model.trim="searchInfo.host"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="hostinfo">
          <el-input
            v-model.trim="searchInfo.hostinfo"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="url">
          <el-input
            v-model.trim="searchInfo.url"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="标题">
          <el-input
            v-model.trim="searchInfo.title"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="favicons">
          <el-input
            v-model.trim="searchInfo.favicons"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="products">
          <el-select
            v-model="searchInfo.products"
            default-first-option
            filterable
            placeholder="请选择"
            allow-create
            remote
            clearable
            multiple
            :reserve-keyword="false"
            :remote-method="productsMethod"
            :loading="productsLoading"
            style="width: 99%"
          >
            <el-option
              v-for="item in productsOptions"
              :key="item"
              :label="item"
              :value="item"
            ></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="protocols">
          <el-input
            v-model.trim="searchInfo.protocols"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="端口是否开放">
          <el-input
            v-model.trim="searchInfo.alive"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="banner">
          <el-input
            v-model.trim="searchInfo.banner"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="状态">
          <el-input
            v-model.trim="searchInfo.status"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="domain_id">
          <el-input
            v-model.trim="searchInfo.domain_id"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="path_ids">
          <el-input
            v-model.trim="searchInfo.path_ids"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="标签">
          <el-input
            v-model.trim="searchInfo.tags"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>
        <el-form-item label="备注">
          <el-input
            v-model.trim="searchInfo.remarks"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
            <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button
              icon="delete"
              size="mini"
              style="margin-left: 10px;"
              :disabled="!multipleSelection.length"
            >删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />

        <el-table-column align="left" label="端口" prop="port" width="120" />

        <el-table-column align="left" label="host" prop="host" width="120" />

        <el-table-column align="left" label="hostinfo" prop="hostinfo" width="120" />

        <el-table-column align="left" label="url" prop="url" width="120" />

        <el-table-column align="left" label="标题" prop="title" width="120" />

        <el-table-column align="left" label="favicons" prop="favicons" width="120" />

        <el-table-column align="left" label="screenshot" prop="screenshot" width="180">
          <template #default="scope">
            <el-image
              style="width: 100%; height: 100%;"
              :src="scope.row.screenshot"
              :preview-src-list="[scope.row.screenshot]"
            />
          </template>
        </el-table-column>

        <el-table-column align="left" label="products" prop="products" width="120" />

        <el-table-column align="left" label="protocols" prop="protocols" width="120" />

        <el-table-column align="left" label="端口是否开放" prop="alive" width="120" />

        <el-table-column align="left" label="banner" prop="banner" width="120" />

        <el-table-column align="left" label="状态" prop="status" width="120" />

        <el-table-column align="left" label="domain_id" prop="domain_id" width="120" />

        <el-table-column align="left" label="path_ids" prop="path_ids" width="120" />

        <el-table-column align="left" label="标签" prop="tags" width="120" />
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column align="left" label="更新时间" prop="update_at" width="200" />
        <el-table-column align="left" fixed="right" label="按钮组" witdh="180">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" @click="dealwithRow(scope.row)">已阅</el-button>
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updatePortInfoFunc(scope.row)"
            >编辑</el-button>
            <el-button type="text" icon="delete" size="mini" @click="deleteRow(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 3, 5, 15, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="端口:">
          <el-input v-model="formData.port" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="host:">
          <el-input v-model="formData.host" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="hostinfo:">
          <el-input v-model="formData.hostinfo" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="url:">
          <el-input v-model="formData.url" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="标题:">
          <el-input v-model="formData.title" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="products:">
          <el-select
            v-model="formData.products"
            default-first-option
            filterable
            placeholder="请选择"
            allow-create
            remote
            clearable
            multiple
            :reserve-keyword="false"
            :remote-method="productsMethod"
            :loading="productsLoading"
            style="width: 99%"
          >
            <el-option v-for="item in productsOptions" :key="item" :label="item" :value="item"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="protocols:">
          <el-input v-model.number="formData.protocols" clearable placeholder="请选择" />
        </el-form-item>

        <el-form-item label="端口是否开放:">
          <el-select v-model="formData.alive" placeholder="请选择" style="width:100%" clearable>
            <el-option
              v-for="(item, key) in intOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="banner:">
          <el-input v-model="formData.banner" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="状态:">
          <el-input v-model="formData.status" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="domain_id:">
          <el-input v-model="formData.domain_id" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="path_ids:">
          <el-input v-model.number="formData.path_ids" clearable placeholder="请选择" />
        </el-form-item>

        <el-form-item label="标签:">
          <el-input v-model="formData.tags" clearable placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:">
          <el-input v-model="formData.remarks" clearable placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'PortInfo'
}
</script>

<script setup>
import {
  createPortInfo,
  deletePortInfo,
  deletePortInfoByIds,
  updatePortInfo,
  findPortInfo,
  getPortInfoList
} from '@/api/port_info'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段


const formData = ref({
  port: "",
  host: "",
  hostinfo: "",
  url: "",
  title: "",
  favicons: "",
  screenshot: "",
  products: [],
  protocols: [],
  alive: undefined,
  banner: "",
  status: "",
  domain_id: "",
  path_ids: [],
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  console.log('搜索', searchInfo.value)
  if (typeof searchInfo.value.products === "string") {
    searchInfo.value.products = searchInfo.value.products.split(",");
  }
  console.log(searchInfo.value.products)
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getPortInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 已处理按钮
const dealwithRow = (row) => {
  console.log(row)
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deletePortInfoFunc(row)
  })
}


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async () => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item._id)
    })
  const res = await deletePortInfoByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updatePortInfoFunc = async (row) => {
  const res = await findPortInfo({ _id: row._id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.report_info
    dialogFormVisible.value = true
  }
}


// 删除行
const deletePortInfoFunc = async (row) => {
  const res = await deletePortInfo({ _id: row._id })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    port: "",
    host: "",
    hostinfo: "",
    url: "",
    title: "",
    favicons: "",
    screenshot: "",
    products: [],
    protocols: [],
    alive: 0,
    banner: "",
    status: "",
    domain_id: "",
    path_ids: [],
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  switch (type.value) {
    case 'create':
      res = await createPortInfo(formData.value)
      break
    case 'update':
      res = await updatePortInfo(formData.value)
      break
    default:
      res = await createPortInfo(formData.value)
      break
  }
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '创建/更改成功'
    })
    closeDialog()
    getTableData()
  }
}


const productsOptions = ref([])
const productsLoading = ref(false)

const productsMethod = async (query) => {
  console.log(query)
  if (query) {
    productsLoading.value = true
    let _portList = await getPortInfoList({ page: 1, pageSize: 20, products: [query] })
    if (_portList.code === 0) {
      setTimeout(() => {
        productsLoading.value = false
        var i;
        for (i in _portList.data.list) {
          var tag
          for (tag of _portList.data.list[i].products) {
            console.log("==" + tag)
            if (tag.includes(query)) {
              productsOptions.value.push(tag)
            }
          }
        }
        productsOptions.value = [...new Set(productsOptions.value)]
        console.log(productsLoading.value)
      }, 200)
    }
  } else {
    productsOptions.value = []
  }
}

const intOptions = [
  {
    value: '0',
    label: '否'
  },
  {
    value: '1',
    label: '是'
  }
]
</script>

<style>
</style>
