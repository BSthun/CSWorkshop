import { Delete, Google } from '@mui/icons-material'
import { Box, Button, Typography } from '@mui/material'
import React from 'react'

const Home = () => {
	return (
		<Box
			sx={{
				display: 'flex',
				width: '100vw',
				height: '100vh',
			}}
		>
			<Box
				sx={{
					flex: 2,
				}}
			>
				<img
					src="https://media.discordapp.net/attachments/756520083903479970/1097917683405701120/frame_001-1.png?width=2228&height=1254"
					alt=""
					width="100%"
					height="100%"
					style={{ objectFit: 'cover' }}
				/>
			</Box>
			<Box
				sx={{
					flex: 1,
					backgroundColor: 'white',
					display: 'flex',
					flexDirection: 'column',
					justifyContent: 'center',
					alignItems: 'center',
				}}
			>
				{/* <Box> */}
				<Typography fontSize={36}>SQL Workshop</Typography>
				<Typography>CSC105 Mentor Session, Backend</Typography>
				<Box
					sx={{
						marginY: 5,
					}}
				>
					<Button variant="outlined" startIcon={<Google />}>
						Sign in with Google
					</Button>
				</Box>

				<Typography>Developed by JIW, Thun, Mixko</Typography>
				{/* </Box> */}
			</Box>
		</Box>
	)
}
export default Home
