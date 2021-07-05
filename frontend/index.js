import React from 'react';
import ReactDOM from 'react-dom';
import {Component} from 'react';



class App extends Component{
// 初始化状态
state = {
  comments:[
    // {id:1, name: 'jace', content:'hhhhhhh'},
    // {id:2, name: 'jack', content:'hhhehehhe'},
    // {id:3, name: 'qurk', content:'hh6666'},
    // {id:4, name: 'ppotr', content:'888888'}
  ],
  userName : '',
  userContent :''
}


//渲染评论列表
renderList(){
  if (this.state.comments.length===0){
    return <div className = "no-comment">暂无留言，快来留言吧~</div>
  }
  else{
    return(
      <ul>
           {this.state.comments.map(item =>(
              // 用唯一key属性遍历
              <li key={item.id}>  
              <h3>留言人：{item.name}</h3>
              <p>评论内容：{item.content}</p>
          </li>
           ))}
        </ul>
    )
  }
}

//处理表单元素值
handleForm = (e) => {
  const {name, value} = e.target
  this.setState(
    {[name]: value}
  )
}

//发表评论
addComment = ()=>{
  const {comments, userName, userContent } = this.state
  //发送log测试
  //console.log(userName, userContent)


  //非空校验
  if (userName.trim() ===''||userContent.trim()===''){
    alert('请输入留言信息！')
    //ruturn停止继续执行
    return
  }

  const newComments = [
    {
      id: Math.random(),
      name:userName,
      content:userContent
    },
    ...comments
  ]

  //console.log(newComments)
  this.setState({
    comments:newComments,
    userName : '',
    userContent :''
  })

}


  render(){
    const {userName, userContent} = this.state
    return(
      <div className = "app">
        <div>
          <input className = "user" type="text" placeholder="请输入留言人" 
            value={userName} name = "userName" onChange={this.handleForm}
          />
          <br/>
          <textarea className="content" id="" cols="30" rows="10" 
          placeholder="请输入留言内容" value={userContent} name = "userContent" onChange = {this.handleForm}
          />
          <br/>
          <button onClick={this.addComment}>提交留言</button>
        </div>
      {this.renderList()}
      </div>

)

  }
}

ReactDOM.render(<App/>, document.getElementById("root"))








