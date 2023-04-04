import React from 'react';
import {
	BrowserRouter,
	Route,
	Routes,
} from 'react-router-dom';
import Facebook from './_facebook/index.jsx';
import Twitter from './_twitter/index.jsx';
import Youtube from './_youtube/index.jsx';
import Preview from './preview/index.jsx';

const AppRouter = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path="/facebook" element={<Facebook />} />
				<Route path="/twitter" element={<Twitter />} />
				<Route path="/youtube" element={<Youtube />} />
				<Route path="/preview" element={<Preview />} />
			</Routes>
		</BrowserRouter>
	);
};

export default AppRouter;
