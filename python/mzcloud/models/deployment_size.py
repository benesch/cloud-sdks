from enum import Enum

class DeploymentSize(str, Enum):
    XS = "XS"
    S = "S"
    M = "M"
    L = "L"
    XL = "XL"

    def __str__(self) -> str:
        return str(self.value)