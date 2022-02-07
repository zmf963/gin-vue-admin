<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">

        
          
            <el-form-item label="POC名称">
              <el-input v-model="searchInfo.poc_name" placeholder="搜索条件" />
            </el-form-item>
          
        
          
            <el-form-item label="POC类型">
              <el-input v-model="searchInfo.poc_type" placeholder="搜索条件" />
            </el-form-item>
          
        
          
            <el-form-item label="POC等级">
              <el-input v-model="searchInfo.poc_level" placeholder="搜索条件" />
            </el-form-item>
          
        
        <el-form-item label="标签">
          <el-input v-model="searchInfo.tags" placeholder="搜索条件" />
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
        <div class="gva-btn-list">
            <el-button size="mini" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button size="mini" type="text" @click="deleteVisible = false">取消</el-button>
                <el-button size="mini" type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" size="mini" style="margin-left: 10px;" :disabled="!multipleSelection.length">删除</el-button>
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

        
          
            <el-table-column align="left" label="POC名称" prop="poc_name" width="120" />
          
        
          
            <el-table-column align="left" label="POC类型" prop="poc_type" width="120" />
          
        
          
            <el-table-column align="left" label="POC等级" prop="poc_level" width="120" />
          
        
        <el-table-column align="left" label="标签" prop="tags" width="120" />
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column align="left" label="更新时间" prop="update_at" width="200" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="text" icon="edit" size="small" class="table-button" @click="updatePocInfoFunc(scope.row)">编辑</el-button>
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

      
        <el-form-item label="POC名称:">
          
            <el-input v-model="formData.poc_name" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="POC类型:">
          
            <el-input v-model="formData.poc_type" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="POC等级:">
          
            <el-input v-model="formData.poc_level" clearable placeholder="请输入" />
          
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
  name: 'PocInfo'
}
</script>

<script setup>
import {
  createPocInfo,
  deletePocInfo,
  deletePocInfoByIds,
  updatePocInfo,
  findPocInfo,
  getPocInfoList
} from '@/api/poc_info'

// 全量引入格式化工具 请按需保留
import {  formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段


const formData = ref({ 
          poc_name: "", 
          
          poc_type: "", 
          
          poc_level: "", 
          
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
const getTableData = async() => {
  const table = await getPocInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deletePocInfoFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
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
      const res = await deletePocInfoByIds({ ids })
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
const updatePocInfoFunc = async(row) => {
    const res = await findPocInfo({ _id: row._id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.repoc_info
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePocInfoFunc = async (row) => {
    const res = await deletePocInfo({ _id: row._id })
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
      
        
          poc_name: "", 
        
      
        
          poc_type: "", 
        
      
        
          poc_level: "", 
        
      
    }
}
// 弹窗确定
const enterDialog = async () => {
      let res
      switch (type.value) {
        case 'create':
          res = await createPocInfo(formData.value)
          break
        case 'update':
          res = await updatePocInfo(formData.value)
          break
        default:
          res = await createPocInfo(formData.value)
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
</script>

<style>
</style>
