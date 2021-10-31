import { post, get } from '../request'

export function getUsers(callback) {
    get("/admin/list").then(res => {
        return callback(res)
    }).catch(err => callback(err))
}

export function delUser(data, callback) {
    post("/admin/del", data).then(res => {
        return callback(res)
    }).catch(err => callback(err))
}