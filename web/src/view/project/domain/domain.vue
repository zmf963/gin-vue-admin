<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="域名">
          <el-input
            v-model.trim="searchInfo.domain"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="IPS">
          <el-input
            v-model.trim="searchInfo.ips"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="HostNams">
          <el-input
            v-model.trim="searchInfo.hostnames"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="操作系统">
          <el-input v-model.trim="searchInfo.os" placeholder="搜索条件" @keyup.enter.native="onSubmit" />
        </el-form-item>

        <el-form-item label="whois信息">
          <el-input
            v-model.trim="searchInfo.whois"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="是否存活">
          <el-input
            v-model.trim="searchInfo.alive"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="cname">
          <el-input
            v-model.trim="searchInfo.cname"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="CDN">
          <el-input
            v-model.trim="searchInfo.cdn"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="cidr">
          <el-input
            v-model.trim="searchInfo.cidr"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="ASN">
          <el-input
            v-model.trim="searchInfo.asn"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="ORG">
          <el-input
            v-model.trim="searchInfo.org"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="地址">
          <el-input
            v-model.trim="searchInfo.addr"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="ISP">
          <el-input
            v-model.trim="searchInfo.isp"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="来源">
          <el-input
            v-model.trim="searchInfo.source"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="TargetName">
          <el-input
            v-model.trim="searchInfo.target_id"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="只看目标已确认" prop="target_id_is_verify">
          <el-switch
            v-model.trim="searchInfo.target_id_is_verify"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>
        <el-form-item label="端口ID列表">
          <el-input
            v-model.trim="searchInfo.port_ids"
            placeholder="搜索条件"
            @keyup.enter.native="onSubmit"
          />
        </el-form-item>

        <el-form-item label="标签">
          <el-input v-model.trim="searchInfo.tags" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model.trim="searchInfo.remarks" placeholder="搜索条件" />
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

        <el-table-column align="left" label="域名" prop="domain" width="160" />

        <el-table-column align="left" label="IPS" prop="ips" width="120" />

        <el-table-column align="left" label="HostNams" prop="hostnames" width="120" />

        <el-table-column align="left" label="操作系统" prop="os" width="120" />

        <el-table-column
          align="left"
          label="whois信息"
          prop="whois"
          width="120"
          :formatter="formatterWhois"
        />

        <el-table-column align="left" label="是否存活" prop="alive" width="120" />

        <el-table-column align="left" label="cname" prop="cname" width="120" />

        <el-table-column align="left" label="CDN" prop="cdn" width="120" />

        <el-table-column align="left" label="cidr" prop="cidr" width="120" />

        <el-table-column align="left" label="ASN" prop="asn" width="120" />

        <el-table-column align="left" label="ORG" prop="org" width="120" />

        <el-table-column align="left" label="地址" prop="addr" width="120" />

        <el-table-column align="left" label="ISP" prop="isp" width="120" />

        <el-table-column align="left" label="来源" prop="source" width="120" />

        <el-table-column align="left" label="TargetName" prop="target_name" width="120" />

        <el-table-column align="left" label="归属目标是否确认" prop="target_id_is_verify" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.target_id_is_verify) }}</template>
        </el-table-column>

        <el-table-column align="left" label="端口ID列表" prop="port_ids" width="120" />

        <el-table-column align="left" label="标签" prop="tags" width="120" />
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column align="left" label="更新时间" prop="update_at" width="200" />
        <el-table-column align="left" fixed="right" label="按钮组" width="180">
          <template #default="scope">
            <el-button type="text" icon="edit" size="small" @click="dealwithRow(scope.row)">处理</el-button>
            <el-button
              type="text"
              icon="edit"
              size="small"
              class="table-button"
              @click="updateDomainFunc(scope.row)"
            >编辑</el-button>
            <el-button type="text" icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
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
        <el-form-item label="域名:">
          <el-input v-model="formData.domain" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="IPS:">
          <el-input v-model.number="formData.ips" clearable placeholder="请选择" />
        </el-form-item>

        <el-form-item label="HostNams:">
          <el-input v-model.number="formData.hostnames" clearable placeholder="请选择" />
        </el-form-item>

        <el-form-item label="操作系统:">
          <el-input v-model="formData.os" clearable placeholder="请输入" />
        </el-form-item>

        <!-- <el-form-item label="whois信息:" /> -->
        <el-form-item label="是否存活:">
          <el-select v-model="formData.alive" placeholder="请选择" style="width:100%" clearable>
            <el-option
              v-for="(item, key) in intOptions"
              :key="key"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="cname:">
          <el-input v-model="formData.cname" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="CDN:">
          <el-select v-model="formData.cdn" placeholder="请选择" style="width:100%" clearable>
            <el-option
              v-for="item in intOptions"
              :key="item.value"
              :label="item.label"
              :value="item.value"
            />
          </el-select>
        </el-form-item>

        <el-form-item label="cidr:">
          <el-input v-model="formData.cidr" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="ASN:">
          <el-input v-model="formData.asn" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="ORG:">
          <el-input v-model="formData.org" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="地址:">
          <el-input v-model="formData.addr" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="ISP:">
          <el-input v-model="formData.isp" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="来源:">
          <el-input v-model="formData.source" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="TargetName:">
          <el-input v-model="formData.target_name" clearable placeholder="请输入" />
        </el-form-item>

        <el-form-item label="归属目标是否确认:">
          <el-switch
            v-model="formData.target_id_is_verify"
            active-color="#13ce66"
            inactive-color="#ff4949"
            active-text="是"
            inactive-text="否"
            clearable
          />
        </el-form-item>

        <el-form-item label="端口ID列表:">
          <el-input v-model.number="formData.port_ids" clearable placeholder="请选择" />
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
  name: 'Domain'
}
</script>

<script setup>
import {
  createDomain,
  deleteDomain,
  deleteDomainByIds,
  updateDomain,
  findDomain,
  getDomainList
} from '@/api/domain'

import { getTargetList } from '@/api/target'
// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'

// 自动化生成的字典（可能为空）以及字段

const formData = ref({
  domain: "",
  ips: [],
  hostnames: [],
  os: "",
  alive: undefined,
  cname: "",
  cdn: undefined,
  cidr: "",
  asn: "",
  org: "",
  addr: "",
  isp: "",
  source: "",
  target_id: "",
  target_id_is_verify: undefined,
  port_ids: [],
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

  if (searchInfo.value.target_id_is_verify === undefined) {
    searchInfo.value.target_id_is_verify = null
  } else {
    searchInfo.value.target_id_is_verify = Boolean(searchInfo.value.target_id_is_verify)
  }

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
  console.log(searchInfo.value.target_id)
  // if (searchInfo.value.target_id !== '') {
  //   const targetData = await getTargetList({ page: 1, pageSize: 1, target_name: searchInfo.value.target_name })
  //   if (targetData.data.list.length > 0) {
  //     searchInfo.value.target_id = targetData.data.list[0].target_id
  //   }
  // }
  // console.log(searchInfo.value.target_id)
  console.log(searchInfo.value.target_id_is_verify)
  const table = await getDomainList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const dealwithRow = async (row) => {
  let res = await findDomain({ _id: row._id })
  console.log(res.data.redomain)
  type.value = 'update'
  if (res.code === 0) {
    if (res.data.redomain.tags instanceof Array) {
      res.data.redomain.tags.push("已阅")
    } else {
      res.data.redomain.tags = ["已阅"]
    }
  }
  console.log(res.data.redomain)
  res = await updateDomain(res.data.redomain)
  console.log("dealwithRow updateDomain ", res.data.redomain)
}


// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteDomainFunc(row)
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
  const res = await deleteDomainByIds({ ids })
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
const updateDomainFunc = async (row) => {
  const res = await findDomain({ _id: row._id })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.redomain
    dialogFormVisible.value = true
  }
}


// 删除行
const deleteDomainFunc = async (row) => {
  const res = await deleteDomain({ _id: row._id })
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
    domain: "",
    ips: [],
    hostnames: [],
    os: "",
    alive: 0,
    cname: "",
    cdn: 0,
    cidr: "",
    asn: "",
    org: "",
    addr: "",
    isp: "",
    source: "",
    target_id: "",
    target_id_is_verify: undefined,
    port_ids: [],
  }
}
// 弹窗确定
const enterDialog = async () => {
  let res
  console.log(formData.value);
  if (typeof formData.value.tags === "string") {
    formData.value.tags = formData.value.tags.split(",");
  }
  switch (type.value) {
    case 'create':
      res = await createDomain(formData.value)
      break
    case 'update':
      res = await updateDomain(formData.value)
      break
    default:
      res = await createDomain(formData.value)
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

const formatterWhois = (row) => {
  return JSON.stringify(row.content);
};
</script>

<style>
</style>
