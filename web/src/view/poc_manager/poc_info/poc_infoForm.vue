<template>
  <div>
    <div class="gva-form-box">
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
  name: 'PocInfo'
}
</script>

<script setup>
import {
  createPocInfo,
  updatePocInfo,
  findPocInfo
} from '@/api/poc_info'

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')



const formData = ref({
        
            poc_name: "", 
          
            poc_type: "", 
          
            poc_level: "", 
          
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findPocInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.repoc_info
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
