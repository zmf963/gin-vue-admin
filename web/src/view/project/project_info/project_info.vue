<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="项目名称">
          <el-input v-model="searchInfo.project_name" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="项目描述">
          <el-input v-model="searchInfo.project_desc" placeholder="搜索条件" />
        </el-form-item>

        <el-form-item label="标签">
          <el-select
            v-model="searchInfo.tags"
            placeholder="请选择"
            multiple
            filterable
            allow-create
            default-first-option
            :reserve-keyword="false"
          >
            <el-option v-for="item in searchInfo.tags" :key="item" :label="item" :value="item">
              <span style="color: #8492a6">{{ item }}</span>
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="searchInfo.remarks" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item>
          <el-button size="mini" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="mini" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <el-row :gutter="20">
        <el-col :span="5">
          <el-card :span="5" class="card-add">
            <el-button icon="plus" style="height: 260px; min-width: 230px;" @click="openDialog">新增</el-button>
          </el-card>
        </el-col>
        <el-col v-for="pro in tableData" :span="5">
          <el-card shadow="hover" class="mycard">
            <template #header>
              <div class="card-header">
                <span class="card-title">
                  <el-icon>
                    <collection />
                  </el-icon>
                  &nbsp;{{ pro.project_name }}
                </span>
                <div calss="card-button">
                  <el-button
                    type="text"
                    icon="edit"
                    size="mini"
                    @click="updateProjectInfoFunc(pro)"
                  >编辑</el-button>
                  <el-button type="text" icon="delete" size="mini" @click="deleteRow(pro)">删除</el-button>
                </div>
              </div>
            </template>
            <el-button size="mini" type="success" @click="onTarget(pro)">目标：{{ pro.target_names }}</el-button>
            <el-button size="mini" type="info" @click="onTask(pro)">任务：{{ pro.task_names }}</el-button>
            <el-button size="mini" type="warning" @click="onDetails(pro)"> 详情</el-button>
            <router-link :to="'/layout/project/target/'"></router-link>
            <br />
            时间：{{ pro.start_time }} ~
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;{{ pro.end_time }}
            <br />
            描述：{{ pro.project_desc }}
            <br />
            <br />标签：
            <el-tag v-for="tag in pro.tags" :key="tag" size="mini" type="info">{{ tag }}&nbsp;</el-tag>
            <br />
            备注：{{ pro.remarks }}
            <br />
            &nbsp;&nbsp;{{ pro.update_at }}
            {{pro}}
          </el-card>
        </el-col>
      </el-row>
    </div>

    <div class="gva-pagination">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[3, 7, 11, 15, 27, 50, 100]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
        <el-form-item label="项目名称:">
          <el-input v-model="formData.project_name" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="项目描述:">
          <el-input v-model="formData.project_desc" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="开始时间:">
          <el-date-picker
            v-model="formData.start_time"
            type="datetime"
            style="width:100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>

        <el-form-item label="结束时间:">
          <el-date-picker
            v-model="formData.end_time"
            type="datetime"
            style="width:100%"
            placeholder="选择日期"
            clearable
          />
        </el-form-item>

        <el-form-item label="目标列表:">
          <el-input v-model="formData.target_ids" clearable placeholder="请选择" />
        </el-form-item>

        <!-- <el-form-item label="任务id列表:">
          <el-input v-model="formData.task_ids" clearable placeholder="请选择" />
        </el-form-item>-->

        <el-form-item label="标签:">
          <el-select
            v-model="formData.tags"
            placeholder="请选择"
            multiple
            filterable
            allow-create
            default-first-option
            :reserve-keyword="false"
          >
            <el-option v-for="item in formData.tags" :key="item" :label="item" :value="item">
              <span style="color: #8492a6">{{ item }}</span>
            </el-option>
          </el-select>
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
  name: 'ProjectInfo'
}
</script>

<script setup>
import {
  createProjectInfo,
  deleteProjectInfo,
  deleteProjectInfoByIds,
  updateProjectInfo,
  findProjectInfo,
  getProjectInfoList
} from '@/api/project_info'

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
//导入router
import router from '@/router'

// 自动化生成的字典（可能为空）以及字段


const formData = ref({
  project_name: "",
  project_desc: "",
  start_time: new Date(),
  end_time: new Date(),
  target_ids: [],
  task_ids: [],
})

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(3)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 3
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
  if (typeof searchInfo.value.tags === 'string') {
    searchInfo.value.tags = searchInfo.value.tags.split(',')
  }
  const table = await getProjectInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteProjectInfoFunc(row)
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
  const res = await deleteProjectInfoByIds({ ids })
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
const updateProjectInfoFunc = async (row) => {
  const res = await findProjectInfo({ _id: row._id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reproject_info
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteProjectInfoFunc = async (row) => {
  const res = await deleteProjectInfo({ _id: row._id })
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
    project_name: "",
    project_desc: "",
    start_time: new Date(),
    end_time: new Date(),
    target_ids: [],
    task_ids: [],
    tags: [],
    remarks: ""
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  console.log(formData.value)
  if (typeof (formData.value.target_ids) === 'string') {
    formData.value.target_ids = formData.value.target_ids.split(',')
  }
  if (typeof (formData.value.task_ids) === 'string') {
    formData.value.task_ids = formData.value.task_ids.split(',')
  }
  switch (type.value) {
    case 'create':
      res = await createProjectInfo(formData.value)
      break
    case 'update':
      res = await updateProjectInfo(formData.value)
      break
    default:
      res = await createProjectInfo(formData.value)
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

const onTarget = (row) => {
  router.push({
    path: '/layout/project/target',
    query: {
      project_id: row._id
    }
  })
}

const onTask = (row) => {
  console.log(row)
  console.log(row._id)
  router.push(
    {
      path: '/layout/project/task',
      query: {
        project_id: row._id
      }
    }
  )
}

const onDetails = (row) => {
  router.push({
    path: '/layout/project/project_details',
    query: {
      project_id: row._id
    }
  })
}

</script>

<style>
.card-header {
  display: flex;
  justify-content: space-between;
  align-items: right;
  font-size: 14px;
}

.card-button {
  justify-content: right;
}

.card-title {
  font-size: 20px;
  color: rgb(60, 46, 104);
}

.card-add {
  margin-top: 25px;
  height: 300px;
  border-radius: 10px;
}

.mycard {
  margin-top: 25px;
  height: 300px;
  border-radius: 10px;
}

.text {
  font-size: 14px;
}

.item {
  margin-bottom: 2px;
}
</style>
