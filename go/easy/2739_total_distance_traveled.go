func distanceTraveled(mainTank int, additionalTank int) int {
    distance := 0
    for mainTank != 0 {
        injection := mainTank / 5
        burned := injection*5
        if mainTank < 5 {
            burned = mainTank
        } 
        mainTank -= burned
        distance += burned*10
        if additionalTank < injection {
            injection = additionalTank
            additionalTank = 0
        } else {
            additionalTank -= injection
        }
        mainTank += injection
    }
    return distance
}
