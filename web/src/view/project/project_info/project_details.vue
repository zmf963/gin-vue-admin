<template>
    <div>
        <div class="gva-table-box">
            <el-row :gutter="24">
                <el-col :span="6">
                    <el-card shadow="never" class="card_top">
                        资产统计
                        1000/1000
                        <el-tag>查看详情</el-tag>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="never" class="card_top">
                        漏洞统计
                        高危
                        中危
                        <el-tag>查看详情</el-tag>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="never" class="card_top">
                        非IT资产
                        账号
                        邮箱
                        文档
                        <el-tag>查看详情</el-tag>
                    </el-card>
                </el-col>
                <el-col :span="6">
                    <el-card shadow="never" class="card_top">
                        其他it资产
                        公众号
                        APP
                        <el-tag>查看详情</el-tag>
                    </el-card>
                </el-col>
            </el-row>

            <el-row :gutter="24">
                <el-col :span="12">目标</el-col>

                <el-col :span="12">文档列表</el-col>
            </el-row>

            <div>
                任务列表
                <div>
                    <el-table style="width: 50%" :data="taskData">
                        <el-table-column align="left" label="任务名称" prop="task_name" width="120" />
                        <el-table-column align="left" label="host" prop="hosts" width="120" />
                        <el-table-column align="left" label="端口" prop="ports" width="120" />
                        <el-table-column align="left" label="关键字" prop="keyword" width="120" />
                        <el-table-column align="left" label="工具列表" prop="tools" width="120" />
                        <el-table-column align="left" label="状态" prop="status" width="120" />
                        <el-table-column align="left" label="标签" prop="tags" width="120" />
                        <el-table-column align="left" label="备注" prop="remarks" width="120" />
                        <el-table-column align="left" label="更新时间" prop="update_at" width="200" />
                    </el-table>
                    <div>
                        <el-pagination
                            layout="total, sizes, prev, pager, next, jumper"
                            :current-page="taskPage"
                            :page-size="taskPageSize"
                            :page-sizes="[7, 3, 5, 15, 30, 50, 100]"
                            :total="taskTotal"
                            @current-change="taskHandleCurrentChange"
                            @size-change="taskHandleSizeChange"
                        />
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'ProjectInfo'
}
</script>

<script setup>


import {
    findProjectInfo,
} from '@/api/project_info'

import {
    getTaskList
} from '@/api/task'



// 全量引入格式化工具 请按需保留
import { formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref } from 'vue'
//导入router
import { useRoute } from "vue-router"


var project_id;

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
    console.log("--------------", project_id)
    const table = await getTaskList({ page: taskPage.value, pageSize: taskPageSize.value, project_id: project_id })
    console.log("========", table)
    if (table.code === 0) {
        taskData.value = table.data.list
        taskTotal.value = table.data.total
        taskPage.value = table.data.page
        taskPageSize.value = table.data.pageSize
    }
}

getTaskData()


const getTargetData = async () => {
    const table = await getTaskList({ page: taskPage.value, pageSize: taskPageSize.value })
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
        path: '/layout/project/target',
        query: {
            project_id: row._id
        }
    })
}

const onTask = (row) => {
    router.push(
        {
            path: '/layout/project/task',
            query: {
                project_id: row._id
            }
        }
    )
}

const onDomain = (row) => {
    router.push({
        path: '/layout/project/domain',
        query: {
            project_id: row._id
        }
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
