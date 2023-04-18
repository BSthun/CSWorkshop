import React from 'react'
import Navbar from '../components/Navbar'
import { Box } from '@mui/material'

const MainLayout: React.FC<any> = ({ children }) => {
	return (
		<>
			<Navbar />
			<Box>{children}</Box>
		</>
	)
}

export default MainLayout
