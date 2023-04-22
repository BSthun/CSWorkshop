import React from 'react'
import Navbar from '../components/Navbar'
import { Box } from '@mui/material'

const MainLayout: React.FC<any> = ({ children }) => {
	return (
		<>
			<Navbar />
			<Box width="100vw" height="calc(100vh - 64px)" marginTop={8}>
				{children}
			</Box>
		</>
	)
}

export default MainLayout
