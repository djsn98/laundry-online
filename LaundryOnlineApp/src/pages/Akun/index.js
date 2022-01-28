import React from 'react'
import { Image, StyleSheet, Text, View } from 'react-native'
import { WARNA_SEKUNDER, WARNA_UTAMA } from '../../utils/constant'
import { IconNoPhotoProfile } from '../../assets';

const Akun = () => {
    const [photoProfile, setPhotoProfile] = React.useState(IconNoPhotoProfile);
    const [username, setUsername] = React.useState('Dennis Jason');
    const [userID, setUserID] = React.useState('@djsn98')

    return (
        <View style={styles.page}>
            <View style={styles.container}>
                <View style={styles.profilePictureContainer}>
                    <Image source={photoProfile} style={styles.profilePicture}/>
                </View>
                <Text style={styles.name}>{username}</Text>
                <Text>{userID}</Text>
            </View>
        </View>
    )
}

export default Akun

const styles = StyleSheet.create({
    page: {
        flex: 1,
        backgroundColor: 'white'
    },
    container: {
        backgroundColor: WARNA_SEKUNDER,
        alignItems: 'center',
        paddingTop: 50,
        paddingBottom: 30,
        borderBottomLeftRadius: 40,
        borderBottomRightRadius: 40,
    },
    profilePictureContainer: {
        width: 125, 
        height: 125,
        borderRadius: 100,
        borderColor: WARNA_UTAMA,
        borderWidth: 1,
        justifyContent: 'center',
        alignItems: 'center'
    },
    profilePicture: {
        width: 120,
        height: 120,
        resizeMode: 'contain',
        borderRadius: 100,
    },
    name: {
        fontSize: 18,
        fontFamily: 'TitilliumWeb-Bold',
        marginTop: 10
    }
})
