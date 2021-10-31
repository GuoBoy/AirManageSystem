import { get } from '../request'

export function getFlightList(callback) {
    get("/v1/flights").then(res => {
        return callback(res)
    }).catch(err => callback(err))
}
