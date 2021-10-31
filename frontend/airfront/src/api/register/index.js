import { postJson } from '../request'
import store from '../../store/index'

export function register(data, callback) {
    postJson("/register", data).then(res => {
        if (res.code === 200) {
            store.commit("updateUserInfo", { u: data.username, t: res.token })
        }
        return callback(res)
    })
}