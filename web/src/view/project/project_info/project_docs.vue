<template>
  <div>
    <div class="gva-table-box">

      <el-pagination layout="total, sizes, prev, pager, next, jumper" :current-page="targetPage"
        :page-size="targetPageSize" :page-sizes="[10, 3, 5, 15, 30, 50, 100]" :total="targetTotal"
        @current-change="targetHandleCurrentChange" @size-change="targetHandleSizeChange" />

      <iframe height=800px width=100% src="http://114.67.244.143:30001/mXMKN-kERzGWmxZG4P_rYQ?both">
      </iframe>

    </div>

  </div>
</template>

<script>
export default {
  name: "ProjectInfo",
}
</script>

<script setup>
import { getTaskList } from "@/api/task";

import { getTargetList } from "@/api/target";

import { findProjectInfo, getProjectInfoList } from "@/api/project_info";

// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, filterDict } from "@/utils/format";
import { ElMessage, ElMessageBox } from "element-plus";
import { ref } from "vue";
//导入router
import { useRoute } from "vue-router";

var project_id

const targetData = ref([])
const targetPage = ref(1)
const targetTotal = ref(0)
const targetPageSize = ref(5)
const DocData = ref({})

const taskData = ref([])
const taskPage = ref(1)
const taskTotal = ref(0)
const taskPageSize = ref(5)

const route = useRoute()

const projectData = ref({})

const getProjectData = async () => {
  project_id = route.query.project_id
  console.log("--------------", project_id)
  const res = await findProjectInfo({ _id: project_id })
  console.log("getProjectData", res)
  if (res.code === 0) {
    projectData.value = res.data.reproject_info.target_ids
  }
  console.log(projectData)
}

getProjectData()

const taskHandleSizeChange = (val) => {
  taskPageSize.value = val
  getTaskData()
}

// 修改页面容量
const taskHandleCurrentChange = (val) => {
  taskPage.value = val
  getTaskData()
}

// 查询
const getTaskData = async () => {
  project_id = route.query.project_id
  const table = await getTaskList({
    page: taskPage.value,
    pageSize: taskPageSize.value,
    project_id: project_id,
  })
  if (table.code === 0) {
    taskData.value = table.data.list
    taskTotal.value = table.data.total
    taskPage.value = table.data.page
    taskPageSize.value = table.data.pageSize
  }
}

getTaskData()

const targetHandleSizeChange = (val) => {
  taskPageSize.value = val
  getTargetData()
}

// 修改页面容量
const targetHandleCurrentChange = (val) => {
  taskPage.value = val
  getTargetData()
}

const getTargetData = async () => {
  project_id = route.query.project_id
  const table = await getTargetList({
    project_id: project_id,
    page: taskPage.value,
    pageSize: taskPageSize.value,
  })
  console.log(table)
  if (table.code === 0) {
    targetData.value = table.data.list
    targetTotal.value = table.data.total
    targetPage.value = table.data.page
    targetPageSize.value = table.data.pageSize
  }
}

getTargetData()

const onTarget = (row) => {
  router.push({
    path: "/layout/project/target",
    query: {
      project_id: row._id,
    },
  })
}

const onTask = (row) => {
  router.push({
    path: "/layout/project/task",
    query: {
      project_id: row._id,
    },
  })
}

const onDomain = (row) => {
  router.push({
    path: "/layout/project/domain",
    query: {
      project_id: row._id,
    },
  })
}
</script>

<style>
.card_top {
  background-color: rgb(252, 248, 248);
  justify-content: space-between;
  height: 160px;
}
</style>
