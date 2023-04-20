import './App.css'
import MainLayout from './layouts/MainLayout'
import Enrollment from './pages/Enrollment'
import Home from './pages/Home'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import Lab from './pages/Lab'
import { ToastContainer, toast } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css'
import './xterm.scss'

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
				fontWeight: 400,
			},
		},
	})

	const router = createBrowserRouter([
		{
			path: '/sqlworkshop/',
			element: <Home />,
		},
		{
			path: '/sqlworkshop/enrollment',
			element: (
				<MainLayout>
					<Enrollment />
				</MainLayout>
			),
		},
		{
			path: '/sqlworkshop/lab/:enrollmentId',
			element: (
				<MainLayout>
					<Lab />
				</MainLayout>
			),
		},
	])

	return (
		<div className="App">
			<ThemeProvider theme={theme}>
				<RouterProvider router={router} />
			</ThemeProvider>
			<ToastContainer />
		</div>
	)
}

export default App
