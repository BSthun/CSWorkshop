import { BatchPrediction } from '@mui/icons-material'
import {
	Box,
	Button,
	CircularProgress,
	Dialog,
	Typography,
} from '@mui/material'
import axios from 'axios'
import React, { useCallback, useRef, useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { BasedResponse } from '../types/APIs/basedResponse'
import { toast } from 'react-toastify'
import { Resizable } from 're-resizable'
import ResizeObserver from 'react-resize-observer'
import { Terminal } from 'xterm'
import { FitAddon } from 'xterm-addon-fit'

const fitAddon = new FitAddon()

function MockComponent() {
	const router = useNavigate()
	const termRef = useRef<Terminal>()
	const initiateTerminal = useCallback(() => {
		setTimeout(() => {
			if (termRef.current) {
				termRef.current.open(document.getElementById('xterm')!)
				return
			}

			termRef.current = new Terminal({
				rows: 20,
				disableStdin: true,
			})
			termRef.current.options = {
				convertEol: true,
				fontFamily: `Inconsolata`,
				fontSize: 18,
				rendererType: 'canvas',
				cursorBlink: true,
			}
			termRef.current.open(document.getElementById('xterm')!)

			// Make the terminal's size and geometry fit the size of #terminal-container
			termRef.current.loadAddon(fitAddon)
			fitAddon.fit()

			// term.onKey(key => {
			//     const char = key.domEvent.key;
			//     if (char === "Enter") {
			//         prompt();

			//     } else if (char === "Backspace") {
			//         term.write("\b \b");
			//     } else {
			//         term.write(char);
			//     }
			// });

			prompt()
		}, 250)

		return () => {
			termRef.current?.dispose()
			termRef.current?.clear()
		}
	}, [])

	const params = useParams()
	const [isLoading, setLoading] = useState(false)
	const [isDialogOpen, setIsDialogOpen] = React.useState(false)
	const websocketRef = useRef<WebSocket>()
	const handleCloseDialog = () => {
		setIsDialogOpen(false)
		// router(0)
	}

	const initializeWebsocket = useCallback(
		(token: string, enrollmentId: string) => {
			if (websocketRef.current) {
				websocketRef.current?.close()
			}
			websocketRef.current = new WebSocket(
				`${
					import.meta.env.VITE_BACKEND_WEBSOCKET_URL
				}/ws/mock?eid=${enrollmentId}&token=${token}`
			)
			websocketRef.current.onopen = (e) => {
				setIsDialogOpen(true)
				initiateTerminal()
			}
			websocketRef.current.onmessage = (e) => {
				if (termRef.current) {
					termRef.current.writeln('\r' + e.data)
				}
			}
		},
		[]
	)

	const onMockDataClicked = useCallback(async () => {
		setLoading(true)
		axios
			.get<BasedResponse<{ token: string }>>(`/api/lab/enroll/mock`, {
				params: {
					enrollmentId: params.enrollmentId,
				},
			})
			.then(({ data }) => {
				const token = data?.data?.token
				toast.success('Success')
				initializeWebsocket(token, String(params.enrollmentId))
			})
			.catch(({ response }) => {
				toast.error(response?.data.message)
			})
			.finally(() => {
				setLoading(false)
			})
	}, [])

	const prompt = () => {
		var shellprompt = '$ '
		termRef.current?.write('\n' + shellprompt)
	}

	// useEffect(() => {
	//     sendJsonMessage({
	//         type: "build.start",
	//         payload: null
	//     });
	// }, [sendJsonMessage]);

	return (
		<>
			<Box
				sx={{
					width: '180px',
					height: '180px',
					borderRadius: 2,
					backgroundColor: 'white',
					boxShadow: '0px 4px 4px rgba(0, 0, 0, 0.25)',
					display: 'flex',
					flexDirection: 'column',
					justifyContent: 'center',
					alignItems: 'center',
					cursor: 'pointer',
				}}
				onClick={onMockDataClicked}
			>
				<Box
					sx={{
						width: '70px',
						height: '90px',
						display: 'flex',
						justifyContent: 'center',
						alignItems: 'center',
					}}
				>
					{isLoading ? (
						<CircularProgress color="info" />
					) : (
						<BatchPrediction
							sx={{
								width: '70px',
								height: '90px',
								color: '#919191',
								mb: 1,
							}}
						/>
					)}
				</Box>
				<Typography>Mock Data</Typography>
			</Box>
			<Dialog open={isDialogOpen} onClose={handleCloseDialog}>
				<Box sx={style}>
					<Typography mb={2}>Generating mock data...</Typography>
					<div
						id="xterm"
						style={{
							height: '100%',
							width: '100%',
							borderRadius: '5px',
							overflow: 'hidden',
						}}
					/>
					<Button
						variant="outlined"
						sx={{
							mt: 3,
							width: '100px',
							height: '40px',
							backgroundColor: '#E1F6FB',
							color: '#0078E7',
							border: 'none',
							textTransform: 'none',
							fontSize: 16,
							':hover': {
								border: 'none',
							},
							alignSelf: 'end',
							borderRadius: 2,
						}}
						onClick={handleCloseDialog}
					>
						Finish
					</Button>
				</Box>
			</Dialog>
		</>
	)
}

const style = {
	position: 'fixed',
	top: '50%',
	left: '50%',
	transform: 'translate(-50%, -50%)',
	display: 'flex',
	flexDirection: 'column',
	width: 720,
	bgcolor: 'white',
	p: 4,
	borderRadius: '8px',
}

export default MockComponent
