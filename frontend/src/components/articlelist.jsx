import React from 'react';
import ReactDOM from 'react-dom';
import axios from 'axios';

const API_ARTICLE_LIST = 'http://127.0.0.1:8080/v1/article';

class ArticleList extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			posts: []
		};
	}

	componentDidMount() {
		axios.get(API_ARTICLE_LIST)
			.then(res => {
				const posts = res.data;
				console.log(posts);
				this.setState({ posts: posts });
			})
			.catch(error => {
				alert("get articles error: " + error);
			});
	}

	render() {
		const articles = this.state.posts.map(post =>
			<li key={post.ArticleId}>
				<h3>{post.Title}</h3><h5>{post.PublishTime}</h5>
				<div>{post.Content}</div>
			</li>
		)
		return (
			<div>
				<h1>Articles</h1>
				<ul>
					{articles}
				</ul>
			</div>
		);
	}
}

export default ArticleList
