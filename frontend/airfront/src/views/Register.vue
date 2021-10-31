<template>
	<div class="register">
		<el-form
			:model="registerData"
			ref="formRef"
			:rules="rules"
			label-width="80px"
		>
			<el-form-item label="用户名" prop="username">
				<el-input
					v-model="registerData.username"
					placeholder="输入用户名"
					clearable
				></el-input> </el-form-item
			><el-form-item label="密码" prop="password">
				<el-input
					v-model="registerData.password"
					placeholder="输入密码"
					type="password"
					clearable
				></el-input> </el-form-item
			><el-form-item label="真实姓名" prop="realname">
				<el-input
					v-model="registerData.realname"
					placeholder="输入真实姓名"
					clearable
				></el-input> </el-form-item
			><el-form-item label="身份证号" prop="idcard">
				<el-input
					v-model="registerData.idcard"
					placeholder="输入身份证号"
					type="password"
					clearable
				></el-input> </el-form-item
			><el-form-item label="电话" prop="phone">
				<el-input
					v-model="registerData.phone"
					placeholder="输入电话号码"
					clearable
				></el-input>
			</el-form-item>
			<el-form-item>
				<div class="btns">
					<router-link to="login"
						><el-button type="primary">登录</el-button></router-link
					>
					<el-button type="primary" @click="onRegister"
						>注册</el-button
					>
				</div>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
import { ref } from "vue";
import { register } from "../api/register";
import { ElMessage } from "element-plus";
import { useRouter } from "vue-router";
export default {
	setup() {
		// 注册提交数据
		const registerData = ref({
			username: "yue",
			password: "yuehongxing",
			realname: "岳红星",
			idcard: "123456789012345678",
			phone: "12345678901",
			issuper: true,
		});
		// 校验规则
		const rules = ref({
			username: [
				{
					required: true,
					message: "请输入用户名",
					trigger: "blur",
				},
				{
					min: 2,
					max: 20,
					message: "长度在2~20个字符",
					trigger: "blur",
				},
			],
			password: [
				{ required: true, message: "请输入密码", trigger: "blur" },
				{
					min: 6,
					max: 12,
					message: "长度在 6 到 12 个字符",
					trigger: "blur",
				},
			],
			realname: [
				{
					required: true,
					message: "输入真实姓名",
					trigger: "blur",
				},
				{
					min: 2,
					max: 20,
					message: "长度在2~20个字符",
					trigger: "blur",
				},
			],
			idcard: [
				{
					validator: (rule, value, callback) => {
						if (value === "") {
							callback(new Error("输入身份证号"));
						} else if (value.length !== 18) {
							callback(new Error("身份证长度不正确"));
						} else {
							callback();
						}
					},
					trigger: "blur",
				},
			],
			phone: [
				{
					validator: (rule, value, callback) => {
						if (value === "") {
							callback(new Error("输入电话号码"));
						} else if (value.length !== 11) {
							callback(new Error("电话号码长度不正确"));
						} else {
							callback();
						}
					},
					trigger: "blur",
				},
			],
		});
		// from ref定义
		const formRef = ref(null);
		const router = useRouter();

		const onRegister = () => {
			formRef.value.validate((v) => {
				if (v) {
					register(registerData.value, (res) => {
						if (res.code === 200) {
							ElMessage.success("注册成功");
							setTimeout(() => {
								router.push("index");
							}, 1000);
						} else ElMessage.error("注册失败, 请检查个人信息");
					});
				} else {
					ElMessage.error("注册失败, 请检查个人信息");
				}
			});
		};

		return {
			registerData,
			rules,
			formRef,
			onRegister,
		};
	},
};
</script>

<style lang="less" scoped>
.register {
	background: url("../assets/bg.jpg") no-repeat;
	position: fixed;
	width: 100%;
	height: 100%;
	.el-form {
		position: absolute;
		left: 50%;
		top: 50%;
		transform: translate(-50%, -50%);
		.btns {
			display: flex;
			justify-content: space-around;
			align-items: center;
		}
	}
}
</style>