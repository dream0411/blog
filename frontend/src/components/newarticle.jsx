import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';

const API_ARTICLE_LIST = 'http://127.0.0.1:8080/v1/article';

class NewArticle extends React.Component {
	constructor(props) {
		super(props);

		this.state = {title: "", content: ""};
		this.submitArticle = this.submitArticle.bind(this);
		this.onTitleChanged = this.onTitleChanged.bind(this);
		this.onContentChanged = this.onContentChanged.bind(this);
	}

	submitArticle(e) {
		e.preventDefault();
		console.log("submit:" + this.state.title + " " + this.state.content);
		this.props.onSubmitArticle(this.state.title, this.state.content);
	}

	onTitleChanged(e) {
		this.setState({title: e.target.value})
		console.log("title:" + this.state.title + " " + this.state.content);
	}

	onContentChanged(e) {
		this.setState({content: e.target.value})
		console.log("content:" + this.state.title + " " + this.state.content);
	}

	render() {
		return <div>
			<form onSubmit={this.submitArticle}>
				<input type="text" placeholder="Type in title..." value={this.state.title} onChange={this.onTitleChanged}/>
				<input type="richtext" value={this.state.content} onChange={this.onContentChanged}/>
				<button type="submit">Submit</button>
			</form>
		</div>
	}
}
export default NewArticle
