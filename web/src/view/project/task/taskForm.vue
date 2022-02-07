<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
   
      
        <el-form-item label="任务名称:">
          
            <el-input v-model="formData.task_name" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="host:">
          
            <el-input v-model="formData.hosts" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="端口:">
          
            <el-input v-model="formData.ports" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="关键字:">
          
            <el-input v-model="formData.keyword" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="工具列表:">
          
            <el-input v-model="formData.tools" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="工具扩展字段:">
          
        </el-form-item>
      
        <el-form-item label="状态:">
          
            <el-input v-model="formData.status" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="项目:">
          
            <el-input v-model="formData.project_id" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="目标:">
          
            <el-input v-model="formData.target_id" clearable placeholder="请输入" />
          
        </el-form-item>
      

        <el-form-item>
          <el-button size="mini" type="primary" @click="save">保存</el-button>
          <el-button size="mini" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Task'
}
</script>

<script setup>
import {
  createTask,
  updateTask,
  findTask
} from '@/api/task'

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')



const formData = ref({
        
            task_name: "", 
          
            hosts: "", 
          
            ports: "", 
          
            keyword: "", 
          
            tools: [], 
          
            status: "", 
          
            project_id: "", 
          
            target_id: "", 
          
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.retask
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }

}

init()
// 保存按钮
const save = async() => {
      let res
      switch (type.value) {
        case 'create':
          res = await createTask(formData.value)
          break
        case 'update':
          res = await updateTask(formData.value)
          break
        default:
          res = await createTask(formData.value)
          break
      }
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '创建/更改成功'
        })
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
