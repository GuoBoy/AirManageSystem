import fs from 'fs'
import path from 'path'

const configPath = path.join(__static, 'default.dll')

const defalutConfig = {
  pdf2iamge: {
    inputpath: '',
    outputpath: '',
    multiple: false,
    single: false
  }
}

const chechConfig = () => {
  fs.stat(configPath, (_, stat) => {
    if (stat && stat.isFile) return
    fs.writeFile(configPath, JSON.stringify(defalutConfig), (err) => {
      if (err) console.log('写入默认配置发生错误', err)
    })
  })
}

const getOneConfig = (module, key, callback) => {
  fs.readFile(configPath, (_, data) => {
    const config = JSON.parse(data.toString())
    try {
      callback(config[module][key])
    } catch {
      callback({})
    }
  })
}

const getConfig = (module, callback) => {
  fs.readFile(configPath, (_, data) => {
    const config = JSON.parse(data.toString())
    try {
      callback(config[module])
    } catch {
      callback('')
    }
  })
}

const updateConfig = (module, value) => {
  fs.readFile(configPath, (_, data) => {
    const config = JSON.parse(data.toString())
    if (config[module] === value) return
    config[module] = value
    fs.writeFile(configPath, JSON.stringify(config), (err) => {
      if (err) console.log('发生错误', err)
    })
  })
}

export {
  chechConfig, getConfig, updateConfig, getOneConfig
}
