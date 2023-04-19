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
			if (termRef.current) termRef.current.reset()

			termRef.current = new Terminal({
				rows: 20,
			})
			termRef.current.options = {
				convertEol: true,
				fontFamily: `Inconsolata`,
				fontSize: 18,

				// fontWeight: 900,
				// back: "#030508",
				rendererType: 'canvas',
				cursorBlink: true,
			}

			if (termRef.current) {
				//Styling
				termRef.current.write('\x1b[31mWelcome to Arttify!\x1b[m\r\n')

				// Load Fit Addon
				termRef.current.loadAddon(fitAddon)

				// Open the terminal in #terminal-container
				termRef.current.open(document.getElementById('xterm')!)

				//Write text inside the terminal
				// term.write(c.magenta("I am ") + c.blue("Blue") + c.red(" and i like it"));

				// Make the terminal's size and geometry fit the size of #terminal-container
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
			}
		}, 250)

		return () => {
			// term.destroy();
			// document.getElementById('terminal')!.remove();
		}
	}, [])

	const params = useParams()
	const [isLoading, setLoading] = useState(false)
	const [isDialogOpen, setIsDialogOpen] = React.useState(false)
	const [log, setLog] = React.useState(
		"Lorem Ipsum is simply dummy text of the printing and typesetting industry. <br /> Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum"
	)
	const websocketRef = useRef<WebSocket>()
	const handleCloseDialog = () => {
		setIsDialogOpen(false)
		router(0)
	}

	const initializeWebsocket = useCallback(
		(token: string, enrollmentId: string) => {
			if (websocketRef.current) {
				websocketRef.current?.close()
			}
			websocketRef.current = new WebSocket(
				`ws://10.4.31.80:3000/ws/mock?eid=${enrollmentId}&token=${token}`
			)
			websocketRef.current.onopen = (e) => {
				setIsDialogOpen(true)
				initiateTerminal()
			}
			websocketRef.current.onmessage = (e) => {
				if (termRef.current) {
					termRef.current.writeln(e.data)
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
			.catch(({ data }) => {
				toast.error(data?.message)
			})
			.finally(() => {
				setLoading(false)
			})
	}, [])

	const prompt = () => {
		var shellprompt = '$ '
		termRef.current?.write('\r' + shellprompt)
	}

	// useEffect(() => {
	//     sendJsonMessage({
	//         type: "build.start",
	//         payload: null
	//     });
	// }, [sendJsonMessage]);

	console.log('rerendered!')

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
			<Dialog open={isDialogOpen} onClose={handleCloseDialog} keepMounted>
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
					<Resizable
						// width={350}
						// height={350}
						style={{
							width: '100%',
							height: '450px',
						}}
					>
						<ResizeObserver
							onResize={(rect) => {
								fitAddon.fit()
								console.log(
									'Resized. New bounds:',
									rect.width,
									'x',
									rect.height
								)
							}}
							onPosition={(rect) => {
								console.log(
									'Moved. New position:',
									rect.left,
									'x',
									rect.top
								)
							}}
						/>
					</Resizable>
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
