<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" label-position="right" label-width="80px">
   
      
        <el-form-item label="域名:">
          
            <el-input v-model="formData.domain" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="IPS:">
          
            <el-input v-model="formData.ips" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="HostNams:">
          
            <el-input v-model="formData.hostnames" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="操作系统:">
          
            <el-input v-model="formData.os" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="whois信息:">
          
        </el-form-item>
      
        <el-form-item label="是否存活:">
          
            
              <el-select v-model="formData.alive" placeholder="请选择" style="width:100%" clearable>
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            
          
        </el-form-item>
      
        <el-form-item label="cname:">
          
            <el-input v-model="formData.cname" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="CDN:">
          
            
              <el-select v-model="formData.cdn" placeholder="请选择" style="width:100%" clearable>
                <el-option v-for="(item,key) in intOptions" :key="key" :label="item.label" :value="item.value" />
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
      
        <el-form-item label="TargetID:">
          
            <el-input v-model="formData.target_id" clearable placeholder="请输入" />
          
        </el-form-item>
      
        <el-form-item label="归属目标是否确认:">
          
            <el-switch v-model="formData.target_id_is_verify" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
          
        </el-form-item>
      
        <el-form-item label="端口ID列表:">
          
            <el-input v-model="formData.port_ids" clearable placeholder="请输入" />
          
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
  name: 'Domain'
}
</script>

<script setup>
import {
  createDomain,
  updateDomain,
  findDomain
} from '@/apidomain'

import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref } from 'vue'
const route = useRoute()
const router = useRouter()
const type = ref('')



const formData = ref({
        
            domain: "", 
          
            ips: [], 
          
            hostnames: [], 
          
            os: "", 
          
            alive:  undefined , 
          
            cname: "", 
          
            cdn:  undefined , 
          
            cidr: "", 
          
            asn: "", 
          
            org: "", 
          
            addr: "", 
          
            isp: "", 
          
            source: "", 
          
            target_id: "", 
          
            target_id_is_verify: false, 
          
            port_ids: [], 
          
        })

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDomain({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redomain
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
      }
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
