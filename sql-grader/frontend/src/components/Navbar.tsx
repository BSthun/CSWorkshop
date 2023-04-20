import { Avatar, Box, Typography } from '@mui/material'
import { useEffect, useState } from 'react'
import { BasedResponse } from '../types/APIs/basedResponse'
import { ProfileAPI } from '../types/APIs/Profile/profile'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

const Navbar = () => {
	const [profile, setProfile] = useState<ProfileAPI>()
	const navigate = useNavigate()

	const fetchState = async () => {
		const profileData = await axios.get<BasedResponse<ProfileAPI>>(
			'/sqlworkshop/api/profile/state'
		)

		setProfile(profileData.data.data)
	}

	useEffect(() => {
		fetchState()
	}, [])

	return (
		<Box
			sx={{
				display: 'flex',
				width: '100vw',
				height: '64px',
				backgroundColor: 'white',
				boxShadow: '0px 4px 8px rgba(0, 0, 0, 0.12)',
				zIndex: 1000,
				position: 'fixed',
				top: 0,
			}}
		>
			<Box
				sx={{
					width: '100%',
					display: 'flex',
					justifyContent: 'space-between',
					alignItems: 'center',
					px: 4,
				}}
			>
				<Typography
					fontSize={24}
					sx={{
						cursor: 'pointer',
					}}
					onClick={() => navigate('/sqlworkshop/enrollment')}
				>
					SQL Playground
				</Typography>
				<Box sx={{ display: 'flex', alignItems: 'center' }}>
					<Avatar
						src={
							profile?.profile.avatar ??
							'https://scontent-sin6-2.xx.fbcdn.net/v/t39.30808-6/338420253_938565190665732_4355582023886486471_n.jpg?_nc_cat=109&ccb=1-7&_nc_sid=09cbfe&_nc_ohc=nlKLB_SKi58AX9LKf5k&_nc_ht=scontent-sin6-2.xx&oh=00_AfAimXgNWQfEMMzfgmktsjMngdF8PEvQe6l6P04OOrCOug&oe=644490A1'
						}
					/>
					<Typography ml={2}>
						{profile?.profile.name ?? 'CSC105 Student'}
					</Typography>
				</Box>
			</Box>
		</Box>
	)
}

export default Navbar
