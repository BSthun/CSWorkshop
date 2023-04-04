import React, {
	useEffect,
	useState,
} from 'react';
import Facebook from '../_facebook/index.jsx';
import Twitter from '../_twitter/index.jsx';
import Youtube from '../_youtube/index.jsx';

const pages = [
	Facebook,
	Twitter,
	Youtube,
];

const pagesMapper = (num) => {
	const Element = pages[num];
	return <Element />;
};

const Preview = () => {
	const [page, setPage] = useState(0);
	
	useEffect(() => {
		const interval = setInterval(() => {
			console.log(1)
			setPage((page) => page + 1);
		}, 10000);
		return () => clearInterval(interval);
	}, []);
	
	return (
		<div>
			{
				pagesMapper(page % 3)
			}
		</div>
	);
};

export default Preview;
