<template>
	<div class="login">
		<el-form
			:model="loginData"
			ref="loginFormRef"
			:rules="rules"
			label-width="80px"
		>
			<el-form-item label="用户名" prop="username">
				<el-input
					v-model="loginData.username"
					placeholder="输入用户名"
					clearable
					class="input-box"
				></el-input>
			</el-form-item>
			<el-form-item label="密  码" prop="password">
				<el-input
					v-model="loginData.password"
					placeholder="输入密码"
					type="password"
					clearable
					class="input-box"
				></el-input>
			</el-form-item>

			<el-form-item>
				<div class="btns">
					<el-button type="primary" class="btn" @click="onLogin"
						>登录</el-button
					>
					<router-link to="register"
						><el-button type="primary" class="btn"
							>注册</el-button
						></router-link
					>
				</div>
			</el-form-item>
		</el-form>
	</div>
</template>

<script>
import { ref } from "vue";
import { ElMessage } from "element-plus";
import { login } from "../api/login";
import { useRouter } from "vue-router";
export default {
	setup() {
		// 登录提交数据
		const loginData = ref({
			username: "",
			password: "",
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
		});
		// from ref定义
		const loginFormRef = ref(null);
		const router = useRouter();
		// 登录
		const onLogin = () => {
			// 验证表单
			loginFormRef.value.validate((v) => {
				if (v) {
					login(loginData.value, (res) => {
						if (res.code === 200) {
							ElMessage.success("登录成功");
							setTimeout(() => {
								router.push("index");
							}, 1000);
						} else {
							ElMessage.warning("登陆失败，请检查用户名或密码");
						}
					});
				} else {
					ElMessage.warning("登陆失败，请检查用户名或密码");
				}
			});
		};

		return {
			loginData,
			rules,
			loginFormRef,
			onLogin,
		};
	},
};
</script>

<style lang="less" scoped>
.login {
	background: url("../assets/bg.jpg") no-repeat;
	position: fixed;
	// top: 0;
	// left: 0;
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