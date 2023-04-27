import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import './App.css'
import Graph from './pages/Graph'

function App() {
	const router = createBrowserRouter([
		{
			path: '/',
			element: <Graph />,
		},
	])

	return (
		<>
			<div>
				<RouterProvider router={router} />
			</div>
		</>
	)
}

export default App
