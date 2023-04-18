import { Avatar, Box, Typography } from '@mui/material'
import React from 'react'

const Navbar = () => {
	return (
		<Box sx={{ width: '100vw', backgroundColor: 'white' }}>
			<Box
				sx={{
					display: 'flex',
					justifyContent: 'space-between',
					alignItems: 'center',
					px: 4,
					py: 1.25,
					boxShadow: '0px 4px 8px rgba(0, 0, 0, 0.12);',
				}}
			>
				<Typography fontSize={24}>SQL Playground</Typography>
				<Box sx={{ display: 'flex', alignItems: 'center' }}>
					<Avatar></Avatar>
					<Typography ml={2}>Apisit Maneerat</Typography>
				</Box>
			</Box>
		</Box>
	)
}

export default Navbar
