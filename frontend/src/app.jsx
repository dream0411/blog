import React from "react"
import ReactDOM from "react-dom"
import ArticleList from "./components/articlelist.jsx"
import NewArticle from "./components/newarticle.jsx"
import axios from 'axios';

class App extends React.Component{
	constructor(props){
		super(props);

		this.handleSubmit = this.handleSubmit.bind(this);
	}

	handleSubmit(title, content) {
		const API_ARTICLE_LIST = 'http://127.0.0.1:8080/v1/article';
		console.log("submit:" + title + " " + content);
		axios.post(API_ARTICLE_LIST, {Title: title, Content: content})
			.then((response) => {
				window.location.reload();
			})
			.catch((error) => {
				alert("submit article error: " + error);
			});
	}
	render(){
		return <div>
				<ArticleList/>
				<hr/>
				<NewArticle onSubmitArticle={this.handleSubmit}/>
			</div>
	}
}

ReactDOM.render(<App/>, document.getElementById("app"))
