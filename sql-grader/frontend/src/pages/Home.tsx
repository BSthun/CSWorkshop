import { Google } from '@mui/icons-material'
import axios from 'axios'
import { GoogleAuthProvider, signInWithPopup } from 'firebase/auth'
import { Box, Button, Typography } from '@mui/material'
import { firebaseAuth } from '../utils/firebase'
import { useNavigate } from 'react-router-dom'

const Home = () => {
	const navigate = useNavigate()

	const processCredential = (result: any) => {
		// No need to do anything with credential. Just want you to see what credential looks like.
		const credential = GoogleAuthProvider.credentialFromResult(result)
		console.log(credential)

		// Get IdToken for verification with backend.
		result.user.getIdToken().then((token: any) => {
			axios
				.post('/api/account/callback', {
					idToken: token,
				})
				.then((response) => {
					if (response.data.success) {
						// Set token from backend to cookie
						// document.cookie = 'token=' + response.data.data.token
						navigate('/enrollment')
					} else {
						alert(response.data.message)
					}
				})
				.catch((err) => {
					if (err.response.status === 500) {
						alert(err.response.data)
					} else {
						alert(err.message)
					}
				})
		})
	}

	const login = () => {
		const provider = new GoogleAuthProvider()
		signInWithPopup(firebaseAuth, provider)
			.then((result) => {
				processCredential(result)
			})
			.catch((error) => {
				alert(error.message)
			})
	}

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
				<Typography fontSize={36}>SQL Workshop</Typography>
				<Typography>CSC105 Mentor Session, Backend</Typography>
				<Box
					sx={{
						marginY: 5,
					}}
				>
					<Button
						variant="outlined"
						startIcon={<Google />}
						onClick={login}
					>
						Sign in with Google
					</Button>
				</Box>

				<Typography>Developed by Thun, Mixko, Jiw</Typography>
			</Box>
		</Box>
	)
}
export default Home
