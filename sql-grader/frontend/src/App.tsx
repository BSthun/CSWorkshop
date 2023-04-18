import './App.css'
import Home from './pages/Home'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

function App() {
	const theme = createTheme({
		palette: {
			primary: {
				main: '#000000',
			},
			secondary: {
				main: '#f44336',
			},
		},
		typography: {
			allVariants: {
				color: 'black',
			},
		},
	})

	const router = createBrowserRouter([
		{
			path: '/',
			element: <Home />,
		},
	])

	return (
		<div className="App">
			<ThemeProvider theme={theme}>
				<RouterProvider router={router} />
			</ThemeProvider>
		</div>
	)
}

export default App
