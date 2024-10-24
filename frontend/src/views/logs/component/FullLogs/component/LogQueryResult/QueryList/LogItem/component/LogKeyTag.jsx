import { Tag } from 'antd'
import React from 'react'
import LogTagDropDown from './LogTagDropdown'
//Key 作为log内容
const LogKeyTag = (props) => {
  const { title, description } = props
  return (
    <div className="flex">
      <div>
        <Tag className="cursor-pointer  text-gray-200">{title}</Tag>
      </div>
      ：
      <LogTagDropDown
        objKey={title}
        value={description}
        children={<span className="break-all hover:underline cursor-pointer">{description}</span>}
      ></LogTagDropDown>
    </div>
  )
}
export default LogKeyTag
