import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Home from './pages/Home'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import { purple } from '@mui/material/colors'

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

	return (
		<div className="App">
			<ThemeProvider theme={theme}>
				<Home />
			</ThemeProvider>
		</div>
	)
}

export default App
