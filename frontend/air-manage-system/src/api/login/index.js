import { post } from '../request'
import store from '../../store/index'

export function login(data, callback) {
    post("/admin/login", data).then(res => {
        if (res.code === 200) {
            store.commit("updateUserInfo", { u: data.username, t: res.token })
        }
        return callback(res)
    })
}