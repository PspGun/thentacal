import * as React from "react"
import
{
  ChakraProvider,
  Box,
  Text,
  Link,
  VStack,
  Code,
  Grid,
  theme,
  Flex,
} from "@chakra-ui/react"
import { ColorModeSwitcher } from "./ColorModeSwitcher"
import { Logo } from "./Logo"
import { Nav } from "./componets/Nav"
import { LoginCard } from "./componets/LoginCard"

export const App = () => (
  <>
    <Nav />
    <Flex
      flexDirection="column"
      width="100wh"
      height="100vh"
      backgroundColor="gray.200"
      justifyContent="center"
      alignItems="center">
      <Box textAlign="center" fontSize="xl">
      </Box>
      <LoginCard />
    </Flex>
  </>
)
